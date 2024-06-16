package handlers

import (
	"net/http"
	"os"

	"github.com/Bipul-Dubey/golang_learn/jwt-project/database"
	"github.com/Bipul-Dubey/golang_learn/jwt-project/models"
	"github.com/Bipul-Dubey/golang_learn/jwt-project/utils"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"github.com/lib/pq"
)

// sign up function
func SignUp(c *gin.Context) {
	var req models.User

	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var err error
	var userID int

	// Hash the password before storing it
	// hashedPassword, err := bcrypt.GenerateFromPassword([]byte(req.Password), bcrypt.DefaultCost)
	// if err != nil {
	// 	c.JSON(http.StatusInternalServerError, gin.H{"error": "Internal server error"})
	// 	return
	// }

	err = database.DB.QueryRow(
		"INSERT INTO users (first_name, last_name, email, password, phone) VALUES ($1, $2, $3, $4, $5) RETURNING id",
		req.FirstName, req.LastName, req.Email, req.Password, req.Phone,
	).Scan(&userID)

	if err != nil {
		if pqErr, ok := err.(*pq.Error); ok && pqErr.Code == "23505" {

			c.JSON(http.StatusNotAcceptable, gin.H{"error": "Email Already exists, please try with another one."})
			return
		}
		c.JSON(http.StatusInternalServerError, gin.H{"error": err})
		return
	}

	c.JSON(http.StatusCreated, gin.H{"message": "User created", "id": userID, "status": http.StatusCreated})
}

func Login(c *gin.Context) {
	var req struct {
		Email    string `json:"email" binding:"required"`
		Password string `json:"password" binding:"required"`
	}
	if err := c.ShouldBindJSON(&req); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	var user models.User
	err := database.DB.QueryRow("SELECT id, first_name, last_name, password, email, phone FROM users WHERE email = $1", req.Email).Scan(
		&user.ID, &user.FirstName, &user.LastName, &user.Password, &user.Email, &user.Phone)
	if err != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
		return
	}

	// Compare the stored hashed password with the provided password
	// if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(req.Password)); err != nil {
	// 	c.JSON(http.StatusUnauthorized, gin.H{"error": "Invalid email or password"})
	// 	return
	// }

	tokenString, err := GenerateToken(user)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Could not generate token"})
		return
	}

	// ========== send token into cookies ==========
	// c.SetSameSite(http.SameSiteLaxMode)
	// c.SetCookie("Authorization", tokenString, 3600*24, "", "", true, true)

	c.JSON(http.StatusOK, gin.H{"userId": user.ID, "token": tokenString,
		"expireAt": utils.GetTokenExpireTime(), "exipreInSec": utils.GetTokenExpireTime().Unix()})
}

func GenerateToken(user models.User) (string, error) {

	token := jwt.NewWithClaims(jwt.SigningMethodHS256, jwt.MapClaims{
		"expireIn": utils.GetTokenExpireTime().Unix(),
		"expireAt": utils.GetTokenExpireTime(),
		"email":    user.Email,
	})

	tokenString, err := token.SignedString([]byte(os.Getenv("SECRET")))
	if err != nil {
		return "", err
	}
	return tokenString, nil
}

func Validate(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "i am logged in",
	})
}

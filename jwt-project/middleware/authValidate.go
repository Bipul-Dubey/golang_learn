package middleware

import (
	"fmt"
	"net/http"
	"os"
	"time"

	"github.com/Bipul-Dubey/golang_learn/jwt-project/database"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
)

// validate authentication
func AuthRequire(c *gin.Context) {
	// get the cookie off request
	// tokenString, err := c.Cookie("Authorization")
	// if err != nil {
	// 	c.AbortWithStatus(http.StatusUnauthorized)
	// }

	// get the tokenString from header
	tokenString := c.GetHeader("Authorization")
	if tokenString == "" {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

	// decode and validate it
	token, _ := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {
		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("unexpected signing method: %v", token.Header["alg"])
		}
		return []byte(os.Getenv("SECRET")), nil
	})

	if claims, ok := token.Claims.(jwt.MapClaims); ok && token.Valid {

		// check the expiration
		if float64(time.Now().Unix()) > claims["expireIn"].(float64) {
			c.AbortWithStatus(http.StatusUnauthorized)
		}

		// find the user with token sub
		var userId int
		err := database.DB.QueryRow("SELECT id FROM users WHERE email = $1", claims["email"]).Scan(&userId)
		if err != nil {
			c.AbortWithStatus(http.StatusUnauthorized)
		}
		// attach the request
		c.Set("user", userId)

		// continue if authentication pass
		c.Next()
	} else {
		c.AbortWithStatus(http.StatusUnauthorized)
	}

}

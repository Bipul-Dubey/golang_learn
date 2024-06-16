package routes

import (
	"github.com/Bipul-Dubey/golang_learn/jwt-project/handlers"
	"github.com/Bipul-Dubey/golang_learn/jwt-project/middleware"
	"github.com/gin-gonic/gin"
)

func Router() {
	router := gin.Default()
	v1 := router.Group("/v1")
	{
		v1.POST("/signup", handlers.SignUp)
		v1.POST("/login", handlers.Login)
		v1.GET("/validate", middleware.AuthRequire, handlers.Validate)
	}

	router.Run(":8080")
}

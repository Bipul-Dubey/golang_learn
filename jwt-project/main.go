package main

import (
	"log"

	"github.com/Bipul-Dubey/golang_learn/jwt-project/database"
	"github.com/Bipul-Dubey/golang_learn/jwt-project/routes"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	// start database connection
	database.DbConnection()

	// starting server
	routes.Router()
}

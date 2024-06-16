package main

import (
	"github.com/Bipul-Dubey/golang_learn/jwt-project/database"
	"github.com/Bipul-Dubey/golang_learn/jwt-project/routes"
)

func main() {
	// start database connection
	database.DbConnection()

	// starting server
	routes.Router()
}

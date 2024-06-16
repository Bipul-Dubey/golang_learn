package database

import (
	"database/sql"
	"fmt"
	"log"

	_ "github.com/lib/pq"
)

var DB *sql.DB

var (
	host     = "localhost"
	port     = "5432"
	user     = "postgres"
	password = "password"
	dbname   = "authentication"
)

func DbConnection() error {
	fmt.Println("Connecting...................")
	// postgres connection config
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// connection
	var err error
	DB, err = sql.Open("postgres", psqlInfo)
	if err != nil {
		log.Fatalf("Failed to connect Database : %s", err)
		return err
	}
	fmt.Println("Connected...................")
	return nil
}

func CloseDB() error {
	fmt.Println("Closing DB connection.............")
	return DB.Close()
}

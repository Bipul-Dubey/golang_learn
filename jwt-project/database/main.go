package database

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func DbConnection() error {
	var err error

	var (
		host     = os.Getenv("HOST")
		port     = os.Getenv("PORT")
		user     = os.Getenv("USER")
		password = os.Getenv("PASSWORD")
		dbname   = os.Getenv("DATABASE")
	)

	fmt.Println("Connecting...................")
	// postgres connection config
	psqlInfo := fmt.Sprintf("host=%s port=%s user=%s password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)

	// connection
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

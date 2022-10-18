package main

import (
	"log"
	"os"

	"github.com/joho/godotenv"
	"github.com/muhangga/internal/config"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	err := godotenv.Load(".env")
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	db, err := config.OpenDB(os.Getenv("MYSQL_URL"), true)
	if err != nil {
		log.Fatal(err)
	}
	defer config.CloseDB(db)
}

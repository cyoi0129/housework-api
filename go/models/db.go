package models

import (
	"database/sql"
	"fmt"
	"os"

	"github.com/joho/godotenv"

	_ "github.com/lib/pq"
)

var DB *sql.DB

func ConnectDatabase() {
	godotenv.Load(".env")
	dbName := os.Getenv("DATABASE_NAME")
	dbUser := os.Getenv("DATABASE_USER")
	dbHost := os.Getenv("DATABASE_HOST")
	dbPassword := os.Getenv("DATABASE_PASSWORD")
	dbPort := 5432

	dsn := fmt.Sprintf("host=%s port=%d user=%s password=%s dbname=%s sslmode=disable",
		dbHost, dbPort, dbUser, dbPassword, dbName,
	)

	db, err := sql.Open("postgres", dsn)
	if err != nil {
		panic(err)
	}
	fmt.Println("DB接続")
	DB = db
}

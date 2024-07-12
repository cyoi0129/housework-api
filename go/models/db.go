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
	dbName := os.Getenv("DBNAME")
	dbUser := os.Getenv("DBUSER")
	dbHost := os.Getenv("DBHOST")
	dbPassword := os.Getenv("DBPASSWORD")
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

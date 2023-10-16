package db

import (
	"database/sql"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func init() {
	if err := godotenv.Load(); err != nil {
		panic("Error loading .env file")
	}
}

func DatabaseConnect() *sql.DB {
	connection := "user=postgres dbname=postgres password=pass host=localhost sslmode=disable"

	// dbUser := os.Getenv("BD_USER")
	// dbPassword := os.Getenv("DB_PASSWORD")
	// dbName := os.Getenv("DB_NAME")
	// dbHost := os.Getenv("DB_HOST")
	// dbMode := os.Getenv("DB_MODE")
	// connection := "user=" + dbUser + " dbname=" + dbName + " password=" + dbPassword + " host=" + dbHost + " sslmode=" + dbMode

	db, err := sql.Open("postgres", connection)
	if err != nil {
		panic(err.Error())
	}
	return db
}

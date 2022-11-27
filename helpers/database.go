package helpers

import (
	"database/sql"
	"log"
	"os"
)

func DbConn() (db *sql.DB) {
	dbDriver := "postgres"
	dbUser := os.Getenv("DATABASE_USERNAME")
	dbPass := os.Getenv("DATABASE_PASSWORD")
	dbName := os.Getenv("DATABASE_NAME")
	dbServer := os.Getenv("DATABASE_HOST")
	dbPort := os.Getenv("DATABASE_PORT")
	log.Println("Database host: " + dbServer)
	log.Println("Database port: " + dbPort)

	db, err := sql.Open(dbDriver, "postgres://"+dbUser+":"+dbPass+"@"+dbServer+":"+dbPort+"/"+dbName+"?sslmode=disable")
	if err != nil {
		panic(err.Error())
	}
	return db
}

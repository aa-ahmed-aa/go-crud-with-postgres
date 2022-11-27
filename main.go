package main

import (
	"log"
	"net/http"
	"os"

	"wallester/handlers"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Server started on: http://localhost:" + os.Getenv("APP_PORT"))
	http.HandleFunc("/", handlers.Index)
	http.HandleFunc("/show", handlers.Show)
	http.HandleFunc("/new", handlers.New)
	http.HandleFunc("/edit", handlers.Edit)
	http.HandleFunc("/insert", handlers.Insert)
	http.HandleFunc("/update", handlers.Update)
	http.HandleFunc("/delete", handlers.Delete)
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil)
}

package main

import (
	"database/sql"
	"html/template"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/joho/godotenv"
	_ "github.com/lib/pq"
)

// Customer struct
type Customer struct {
	Id        int
	FirstName string
	LastName  string
	Birthday  string
	Gender    int
	Email     string
	Address   string
}

func dbConn() (db *sql.DB) {
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

var tmpl = template.Must(template.ParseGlob("templates/*"))

// Index handler
func Index(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	selDB, err := db.Query("SELECT * FROM customers ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	customer := Customer{}
	res := []Customer{}

	for selDB.Next() {
		var id, gender int
		var first_name, last_name, email, address, birthday string
		err := selDB.Scan(&id, &first_name, &last_name, &birthday, &gender, &email, &address)
		if err != nil {
			panic(err.Error())
		}

		customer.Id = id
		customer.FirstName = first_name
		customer.LastName = last_name
		customer.Email = email
		customer.Birthday = strings.Split(birthday, "T")[0]
		customer.Address = address
		customer.Gender = gender

		res = append(res, customer)
	}
	tmpl.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

// Show handler
func Show(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM customers WHERE id=$1", nId)
	if err != nil {
		panic(err.Error())
	}

	customer := Customer{}

	for selDB.Next() {
		var id, gender int
		var first_name, last_name, email, address, birthday string
		err := selDB.Scan(&id, &first_name, &last_name, &birthday, &gender, &email, &address)
		if err != nil {
			panic(err.Error())
		}

		customer.Id = id
		customer.FirstName = first_name
		customer.LastName = last_name
		customer.Email = email
		customer.Birthday = strings.Split(birthday, "T")[0]
		customer.Address = address
		customer.Gender = gender
	}
	tmpl.ExecuteTemplate(w, "Show", customer)
	defer db.Close()
}

func New(w http.ResponseWriter, r *http.Request) {
	tmpl.ExecuteTemplate(w, "New", nil)
}

func Edit(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM customers WHERE id=$1", nId)
	if err != nil {
		panic(err.Error())
	}

	customer := Customer{}

	for selDB.Next() {
		var id, gender int
		var first_name, last_name, email, address, birthday string
		err := selDB.Scan(&id, &first_name, &last_name, &birthday, &gender, &email, &address)
		if err != nil {
			panic(err.Error())
		}

		customer.Id = id
		customer.FirstName = first_name
		customer.LastName = last_name
		customer.Email = email
		customer.Birthday = strings.Split(birthday, "T")[0]
		customer.Address = address
		customer.Gender = gender
	}

	tmpl.ExecuteTemplate(w, "Edit", customer)
	defer db.Close()
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		email := r.FormValue("email")
		birthday := r.FormValue("birthday")
		address := r.FormValue("address")
		gender := r.FormValue("gender")
		insForm, err := db.Prepare("INSERT INTO customers (first_name, last_name, birthday, gender, email, address) VALUES ($1, $2, $3, $4, $5, $6) RETURNING id")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(first_name, last_name, email, birthday, address, gender)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	if r.Method == "POST" {
		id := r.FormValue("id")
		first_name := r.FormValue("first_name")
		last_name := r.FormValue("last_name")
		email := r.FormValue("email")
		birthday := r.FormValue("birthday")
		address := r.FormValue("address")
		gender := r.FormValue("gender")
		insForm, err := db.Prepare("UPDATE customers SET first_name=$1, last_name=$2, birthday=$3, gender=$4, email=$5, address=$6 WHERE id=$7")
		if err != nil {
			panic(err.Error())
		}
		insForm.Exec(first_name, last_name, birthday, gender, email, address, id)
	}
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func Delete(w http.ResponseWriter, r *http.Request) {
	db := dbConn()
	customer := r.URL.Query().Get("id")
	delForm, err := db.Prepare("DELETE FROM customers WHERE id=$1")
	if err != nil {
		panic(err.Error())
	}
	delForm.Exec(customer)
	log.Println("DELETE " + customer)
	defer db.Close()
	http.Redirect(w, r, "/", 301)
}

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}
	log.Println("Server started on: http://localhost:" + os.Getenv("APP_PORT"))
	http.HandleFunc("/", Index)
	http.HandleFunc("/show", Show)
	http.HandleFunc("/new", New)
	http.HandleFunc("/edit", Edit)
	http.HandleFunc("/insert", Insert)
	http.HandleFunc("/update", Update)
	http.HandleFunc("/delete", Delete)
	http.ListenAndServe(":"+os.Getenv("APP_PORT"), nil)
}

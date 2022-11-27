package handlers

import (
	"net/http"
	"wallester/helpers"
)

func New(w http.ResponseWriter, r *http.Request) {
	Template.ExecuteTemplate(w, "New", nil)
}

func Insert(w http.ResponseWriter, r *http.Request) {
	db := helpers.DbConn()
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

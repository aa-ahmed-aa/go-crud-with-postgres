package handlers

import (
	"net/http"
	"strings"
	"wallester/helpers"
	"wallester/models"
)

func Edit(w http.ResponseWriter, r *http.Request) {
	db := helpers.DbConn()
	nId := r.URL.Query().Get("id")
	selDB, err := db.Query("SELECT * FROM customers WHERE id=$1", nId)
	if err != nil {
		panic(err.Error())
	}

	customer := models.Customer{}

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

	Template.ExecuteTemplate(w, "Edit", customer)
	defer db.Close()
}

func Update(w http.ResponseWriter, r *http.Request) {
	db := helpers.DbConn()
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

package handlers

import (
	"html/template"
	"net/http"
	"strings"
	"wallester/helpers"
	"wallester/models"
)

var Template = template.Must(template.ParseGlob("templates/*"))

func Index(w http.ResponseWriter, r *http.Request) {
	db := helpers.DbConn()
	selDB, err := db.Query("SELECT * FROM customers ORDER BY id DESC")
	if err != nil {
		panic(err.Error())
	}

	customer := models.Customer{}
	res := []models.Customer{}

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
	Template.ExecuteTemplate(w, "Index", res)
	defer db.Close()
}

package handlers

import (
	"net/http"
	"strings"
	"wallester/helpers"
	"wallester/models"
)

// Show handler
func Show(w http.ResponseWriter, r *http.Request) {
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
	Template.ExecuteTemplate(w, "Show", customer)
	defer db.Close()
}

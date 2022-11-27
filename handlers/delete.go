package handlers

import (
	"log"
	"net/http"
	"wallester/helpers"
)

func Delete(w http.ResponseWriter, r *http.Request) {
	db := helpers.DbConn()
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

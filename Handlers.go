package main

import (
	"fmt"
	"html"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
)


func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func ListAllEmployees(w http.ResponseWriter, r *http.Request) {
	employees := Employees{
		Employee{ID: "SCOTT"},
		Employee{ID: "RUSSELL"},
	}

	json.NewEncoder(w).Encode(employees)
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]
	fmt.Fprintln(w, "employee searching: ", id)
}

package main

import (
	"fmt"
	"html"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
	"io"
	"strconv"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func getEmployees(w http.ResponseWriter, r *http.Request) {
	employees := RepoFindEmployees()

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(employees); err != nil {
		panic(err)
	}
}

func getEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	employee := RepoFindEmployee(id)

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(employee); err != nil {
		panic(err);
	}
}

func createEmployee(w http.ResponseWriter, r *http.Request) {
	var employee Employee
	body, error := ioutil.ReadAll(io.LimitReader(r.Body, 121));

	if error != nil {
		panic(error)
	}

	if error := r.Body.Close(); error != nil {
		panic(error)
	}

	if error := json.Unmarshal(body, &employee); error != nil {
		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if error := json.NewEncoder(w).Encode(error); error != nil {
			panic(error)
		}
	}

	t := RepoCreateEmployee(employee)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err);
	}
}

func updateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, errorParam := strconv.Atoi(vars["id"])

	if errorParam != nil {
		panic(errorParam)
	}

	var employee Employee
	body, errorBody := ioutil.ReadAll(io.LimitReader(r.Body, 121))

	if errorBody != nil {
		panic(errorBody)
	}

	if errorBody := r.Body.Close(); errorBody != nil {
		panic(errorBody)
	}

	if error := json.Unmarshal(body, &employee); error != nil {
		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity
		if error := json.NewEncoder(w).Encode(error); error != nil {
			panic(error)
		}
	}

	t := RepoUpdateEmployee(id , employee)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err);
	}
}

func deleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if error := RepoDestroyEmployee(id); error != nil {
		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity --> 뭐라고 정의해야하나?
		panic(error)
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK) //성공은 뭐라고 정의 하나??
}
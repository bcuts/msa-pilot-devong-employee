package handler

import (
	"fmt"
	"html"
	"net/http"
	"github.com/gorilla/mux"
	"encoding/json"
	"io/ioutil"
	"io"
	"strconv"
	"act-msa-pilot-devong-employee/repository"
	"act-msa-pilot-devong-employee/model"
)

func Index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, %q", html.EscapeString(r.URL.Path))
}

func GetEmployees(w http.ResponseWriter, r *http.Request) {
	employees := repository.RepoFindEmployees()

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(employees); err != nil {
		panic(err)
	}
}

func GetEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	employee := repository.RepoFindEmployee(id)

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)

	if err := json.NewEncoder(w).Encode(employee); err != nil {
		panic(err);
	}
}

func CreateEmployee(w http.ResponseWriter, r *http.Request) {
	var employee model.Employee
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

	t := repository.RepoCreateEmployee(employee)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusCreated)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err);
	}
}

func UpdateEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, errorParam := strconv.Atoi(vars["id"])

	if errorParam != nil {
		panic(errorParam)
	}

	var employee model.Employee
	body, errorBody := ioutil.ReadAll(io.LimitReader(r.Body, 121))

	if errorBody != nil {
		panic(errorBody)
	}

	if errorBody := r.Body.Close(); errorBody != nil {
		panic(errorBody)
	}

	if error := json.Unmarshal(body, &employee); error != nil {
		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		w.WriteHeader(http.StatusUnprocessableEntity) // unprocessable entity
		if error := json.NewEncoder(w).Encode(error); error != nil {
			panic(error)
		}
	}

	employee.ID = id
	t := repository.RepoUpdateEmployee(employee)
	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK)
	if err := json.NewEncoder(w).Encode(t); err != nil {
		panic(err);
	}
}

func DeleteEmployee(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, _ := strconv.Atoi(vars["id"])

	if error := repository.RepoDestroyEmployee(id); error != nil {
		w.Header().Set("Content-type", "application/json;charset=UTF-8")
		w.WriteHeader(422) // unprocessable entity --> 뭐라고 정의해야하나?
		panic(error)
	}

	w.Header().Set("Content-type", "application/json;charset=UTF-8")
	w.WriteHeader(http.StatusOK) //성공은 뭐라고 정의 하나??
}
package main

import (
	"net/http"
)

type Route struct {
	Name        string
	Method      string
	Pattern     string
	HandlerFunc http.HandlerFunc
}

type Routes []Route

var routes = Routes{
	Route{
		"Index",
		"GET",
		"/",
		Index,
	},

	Route{
		"EmployeeList",
		"GET",
		"/employees",
		getEmployees,
	},
	Route{
		"ShowEmployee",
		"GET",
		"/employees/{id}",
		getEmployee,
	},
	Route{
		"EmployCreate",
		"POST",
		"/employees",
		createEmployee,
	},
	Route{
		"UpdateEmployee",
		"PUT",
		"/employees/{id}",
		updateEmployee,
	},
	Route{
		"DeleteEmployee",
		"DELETE",
		"/employees/{id}",
		deleteEmployee,
	},
}
package router

import (
	"net/http"
	"act-msa-pilot-devong-employee/handler"
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
		handler.Index,
	},
	Route{
		"EmployeeList",
		"GET",
		"/employees",
		handler.GetEmployees,
	},
	Route{
		"ShowEmployee",
		"GET",
		"/employees/{id}",
		handler.GetEmployee,
	},
	Route{
		"EmployCreate",
		"POST",
		"/employees",
		handler.CreateEmployee,
	},
	Route{
		"UpdateEmployee",
		"PUT",
		"/employees/{id}",
		handler.UpdateEmployee,
	},
	Route{
		"DeleteEmployee",
		"DELETE",
		"/employees/{id}",
		handler.DeleteEmployee,
	},
}
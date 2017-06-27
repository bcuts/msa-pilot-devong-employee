package router

import (
	"github.com/gorilla/mux"
	"act-msa-pilot-devong-employee/middle"
)

func NewRouter() *mux.Router {

	router := mux.NewRouter().StrictSlash(true)
	for _, route := range routes {
		router.Methods(route.Method).
			Path(route.Pattern).
			Name(route.Name).
			Handler(middle.Logger(route.HandlerFunc, route.Name))
	}
	return router
}

package routes

import (
	"net/http"
	"webapp/src/middlewares"

	"github.com/gorilla/mux"
)

type Route struct {
	URI                    string
	Method                 string
	Function               func(http.ResponseWriter, *http.Request)
	RequiresAuthentication bool
}

func Configuration(router *mux.Router) *mux.Router {
	routes := loginRoutes
	routes = append(routes, userRoutes...)
	routes = append(routes, homeRoute)

	for _, route := range routes {
		if route.RequiresAuthentication {
			router.HandleFunc(route.URI, middlewares.Logger(middlewares.Auth(route.Function))).Methods(route.Method)
		} else {
			router.HandleFunc(route.URI,
				middlewares.Logger(route.Function),
			).Methods(route.Method)
		}
	}

	fileServer := http.FileServer(http.Dir("./assets/"))
	router.PathPrefix("/assets/").Handler(http.StripPrefix("/assets/", fileServer))

	return router
}

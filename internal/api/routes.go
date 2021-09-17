package api

import "net/http"

func (a *Api) ConfigureRoutes() {
	a.Router.Handle(
		"/version",
		http.HandlerFunc(a.HandleVersion)).
		Methods("GET")

	a.Router.Handle(
		"/users",
		a.withMFA(http.HandlerFunc(a.HandleRegister))).
		Methods("POST")

	a.Router.Handle(
		"/users/{phone}",
		a.withMFA(http.HandlerFunc(a.HandleReadProfile))).
		Methods("GET")

	a.Router.Handle(
		"/users/{phone}",
		a.withMFA(http.HandlerFunc(a.HandleUnregister))).
		Methods("DELETE")
}

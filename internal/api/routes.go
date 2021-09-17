package api

import "net/http"

func (a *Api) ConfigureRoutes() {
	a.Router.Handle("/users", a.withMFA(http.HandlerFunc(a.Register))).Methods("POST")
	a.Router.Handle("/users/{phone}", a.withMFA(http.HandlerFunc(a.GetProfile))).Methods("GET")
	a.Router.Handle("/users/{phone}", a.withMFA(http.HandlerFunc(a.Unregister))).Methods("DELETE")
}

package api

import (
	"net/http"
)

func (a *Api) Register(w http.ResponseWriter, r *http.Request)   {}
func (a *Api) Unregister(w http.ResponseWriter, r *http.Request) {}
func (a *Api) GetProfile(w http.ResponseWriter, r *http.Request) {}

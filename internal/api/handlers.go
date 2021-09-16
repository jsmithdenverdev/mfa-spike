package api

import (
	"encoding/json"
	"mfaspike/internal/domain/commands"
	"mfaspike/internal/domain/queries"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *Api) Register(w http.ResponseWriter, r *http.Request) {
	type registerRequest struct {
		Phone    string `json:"phone"`
		Name     string `json:"name"`
		Timezone string `json:"timezone"`
	}

	req := registerRequest{}
	dec := json.NewDecoder(r.Body)
	err := dec.Decode(&req)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	_, err = a.Commands.CreateUser.Handle(commands.CreateUserRequest{
		Name:     req.Name,
		Contact:  req.Phone,
		Timezone: req.Timezone,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusCreated)
}

func (a *Api) Unregister(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phone := vars["phone"]

	if phone == "" {
		http.Error(w, "no phone supplied", http.StatusBadRequest)
		return
	}

	_, err := a.Commands.DeleteUser.Handle(commands.DeleteUserRequest{
		Contact: phone,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

func (a *Api) GetProfile(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	phone := vars["phone"]

	if phone == "" {
		http.Error(w, "no phone supplied", http.StatusBadRequest)
		return
	}

	result, err := a.Queries.ReadUser.Handle(queries.ReadUserRequest{
		Contact: phone,
	})

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	enc := json.NewEncoder(w)
	err = enc.Encode(result.User)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

	w.WriteHeader(http.StatusOK)
}

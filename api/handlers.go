package api

import (
	"encoding/json"
	"mfaspike/commands"
	"mfaspike/queries"
	"net/http"

	"github.com/gorilla/mux"
)

func (a *Api) HandleRegister(w http.ResponseWriter, r *http.Request) {
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

func (a *Api) HandleUnregister(w http.ResponseWriter, r *http.Request) {
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
}

func (a *Api) HandleReadProfile(w http.ResponseWriter, r *http.Request) {
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

	w.Header().Add("Content-Type", "application/json")

	enc := json.NewEncoder(w)
	err = enc.Encode(result.User)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

func (a *Api) HandleVersion(w http.ResponseWriter, r *http.Request) {
	type versionResponse struct {
		Version string `json:"version"`
	}

	w.Header().Add("Content-Type", "application/json")

	res := versionResponse{Version: "1.0.0"}
	enc := json.NewEncoder(w)
	err := enc.Encode(&res)

	if err != nil {
		http.Error(w, err.Error(), http.StatusInternalServerError)
		return
	}

}

package handlers

import (
	"encoding/json"
	"net/http"
)

type CreateUser struct{}

func NewCreateUser() *CreateUser {
	return &CreateUser{}
}

func (h *CreateUser) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "error decoding JSON", http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{"message": "user created", "data": data}
	writeJSON(w, response, http.StatusCreated)
}

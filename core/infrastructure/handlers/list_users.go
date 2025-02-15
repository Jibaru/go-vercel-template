package handlers

import "net/http"

type ListUsers struct{}

func NewListUsers() *ListUsers {
	return &ListUsers{}
}

func (h *ListUsers) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"users": []map[string]interface{}{
			{"id": 1, "name": "Alice"},
			{"id": 2, "name": "Bob"},
		},
	}
	writeJSON(w, response, http.StatusOK)
}

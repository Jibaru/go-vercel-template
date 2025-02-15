package handlers

import "net/http"

type Home struct {
}

func NewHome() *Home {
	return &Home{}
}

func (h *Home) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	writeJSON(w, map[string]string{"message": "Hello"}, http.StatusOK)
}

package handler

import (
	"net/http"

	"github.com/jibaru/go-vercel-template/core/infrastructure/server"
)

var mux = server.New()

func Handler(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}

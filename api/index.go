package handler

import (
	"net/http"

	"github.com/jibaru/go-vercel-template/core/router"
)

var mux = router.Server()

func Handler(w http.ResponseWriter, r *http.Request) {
	mux.ServeHTTP(w, r)
}

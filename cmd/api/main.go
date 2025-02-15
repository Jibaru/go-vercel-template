package main

import (
	"net/http"

	"github.com/jibaru/go-vercel-template/core/infrastructure/server"
)

func main() {
	s := server.New()

	if err := http.ListenAndServe(":8080", s); err != nil {
		panic(err)
	}
}

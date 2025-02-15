package server

import (
	"net/http"

	"github.com/jibaru/go-vercel-template/core/infrastructure/handlers"
)

type Route struct {
	Method  string
	Path    string
	Handler http.Handler
}

func routes() []Route {
	homeHandler := handlers.NewHome()
	listUsersHandler := handlers.NewListUsers()
	createUserHandler := handlers.NewCreateUser()

	return []Route{
		{"GET", "/api", homeHandler},
		{"GET", "/api/users", listUsersHandler},
		{"POST", "/api/users", createUserHandler},
	}
}

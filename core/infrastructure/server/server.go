package server

import (
	"fmt"
	"net/http"
)

func New() *http.ServeMux {
	mux := http.NewServeMux()
	routes := routes()

	for _, route := range routes {
		mux.Handle(fmt.Sprintf("%v %v", route.Method, route.Path), route.Handler)
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	return mux
}

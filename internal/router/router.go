package router

import (
	"encoding/json"
	"net/http"
)

type Route struct {
	Method  string
	Path    string
	Handler http.HandlerFunc
}

func Server() *http.ServeMux {
	var mux = http.NewServeMux()

	var routes = []Route{
		{"GET", "/api", homeHandler},
		{"GET", "/api/users", usersHandler},
		{"POST", "/api/users", createUserHandler},
		{"GET", "/api/products", productsHandler},
		{"POST", "/api/products", createProductHandler},
	}

	for _, route := range routes {
		mux.HandleFunc(route.Path, methodHandler(route.Method, route.Handler))
	}

	mux.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		http.NotFound(w, r)
	})

	return mux
}

func methodHandler(method string, handler http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		if r.Method != method {
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
			return
		}
		handler(w, r)
	}
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]string{"message": "Hello"}
	writeJSON(w, response, http.StatusOK)
}

func usersHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"users": []map[string]interface{}{
			{"id": 1, "name": "Alice"},
			{"id": 2, "name": "Bob"},
		},
	}
	writeJSON(w, response, http.StatusOK)
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]string
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "error decoding JSON", http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{"message": "user created", "data": data}
	writeJSON(w, response, http.StatusCreated)
}

func productsHandler(w http.ResponseWriter, r *http.Request) {
	response := map[string]interface{}{
		"products": []map[string]interface{}{
			{"id": 1, "name": "Laptop", "price": 1200},
			{"id": 2, "name": "Mouse", "price": 20},
		},
	}
	writeJSON(w, response, http.StatusOK)
}

func createProductHandler(w http.ResponseWriter, r *http.Request) {
	var data map[string]interface{}
	if err := json.NewDecoder(r.Body).Decode(&data); err != nil {
		http.Error(w, "error decoding JSON", http.StatusBadRequest)
		return
	}
	response := map[string]interface{}{"message": "product created", "data": data}
	writeJSON(w, response, http.StatusCreated)
}

func writeJSON(w http.ResponseWriter, data interface{}, status int) {
	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(status)
	json.NewEncoder(w).Encode(data)
}

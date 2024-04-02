package main

import (
	"net/http"
)

func main() {

    // Create a new request multiplexer
    // Take incoming requests and dispatch them to the matching handlers
    mux := http.NewServeMux()
    // Register the routes and handlers
    mux.Handle("/", &homeHandler{})
    mux.Handle("/recipes", &RecipesHandler{})
    mux.Handle("/recipes/", &RecipesHandler{})
    // Run the server
    http.ListenAndServe(":8080", mux)
}

type RecipesHandler struct{}
type homeHandler struct{}


func (h *RecipesHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
    w.Write([]byte("This is my recipe page"))
}
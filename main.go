package main

import (
	"fmt"
	"net/http"

	"github.com/go-chi/chi/v5"
	"github.com/go-chi/chi/v5/middleware"
)

func main() {
	// create a chi router, chi makes routing easy
	router := chi.NewRouter()

	// the use method allows use to have a function wrap the handler function before executing (middleware)
	router.Use(middleware.Logger)

	// a router just needs a path and a handler function
	router.Get("/hello", basicHandler)

	// still use the http.Server struct but the handler now uses the rotuer
	server := &http.Server{
		Addr:    ":3000",
		Handler: router,
	}

	err := server.ListenAndServe()
	if err != nil {
		fmt.Println("Server failed to run", err)
	}

}

func basicHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Hello World"))
}

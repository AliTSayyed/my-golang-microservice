package application

import (
	"context"
	"fmt"
	"net/http"
)

// this is the "instance variable of what our application is"
type App struct {
	router http.Handler
}

// create a New function that instantiates the app variable and make it a pointer to that variable
func New() *App {
	app := &App{
		router: loadRoutes(),
	}
	return app
}

// takes a pointer to that variable and then starts the server
func (a *App) Start(ctx context.Context) error {
	server := &http.Server{
		Addr:    ":3000",
		Handler: a.router,
	}
	err := server.ListenAndServe()
	if err != nil {
		return fmt.Errorf("failed to start server: %w", err)
	}
	return nil
}

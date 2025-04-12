package main

import (
	"context"
	"fmt"

	"github.com/alitsayyed/my-golang-microservice/application"
)

func main() {
	// create a new instance of out app
	app := application.New()
	err := app.Start(context.TODO())
	if err != nil {
		fmt.Println("failed to start app:", err)
	}
}

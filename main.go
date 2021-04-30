package main

import (
	"context"
	"fmt"
	"goproj/BusinessLogic"
	"goproj/Route"
	"time"
)

func main() {
	fmt.Println("Starting the application...")
	ctx, _ := context.WithTimeout(context.Background(), 10*time.Second)
	BusinessLogic.CreateConnection(ctx)
	Route.RegisterRoutes()
}
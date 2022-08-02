package main

import (
	"github.com/xvbnm48/go-microservice-udemy/app"
	"github.com/xvbnm48/go-microservice-udemy/logger"
)

func main() {
	// log.Println("Starting the application...")
	logger.Info("Starting the application at port 8080")
	app.Start()
}

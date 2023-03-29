package main

import (
	"github.com/cassiusbessa/create-text/handlers"
	"github.com/cassiusbessa/create-text/logs"
	"github.com/cassiusbessa/create-text/repositories"
	"github.com/sirupsen/logrus"
)

var file = logs.Init()

func main() {
	defer file.Close()
	r := handlers.Router()
	repositories.Repo.Ping()
	r.PUT("/texts/:company/:section", handlers.Handle)
	r.StaticFile("/logs", "./logs/logs.log")
	if err := r.Run(":8080"); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}

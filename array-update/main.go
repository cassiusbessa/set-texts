package main

import (
	"github.com/cassiusbessa/db-texts/handlers"
	"github.com/cassiusbessa/db-texts/logs"
	"github.com/cassiusbessa/db-texts/repositories"
	"github.com/sirupsen/logrus"
)

var file = logs.Init()

func main() {
	defer file.Close()
	r := handlers.Router()
	repositories.Repo.Ping()
	r.PATCH("/texts/:company/:section/:id", handlers.Handle)
	r.StaticFile("/logs", "./logs/logs.log")
	if err := r.Run(":8080"); err != nil {
		logrus.Fatalf("Failed to start server: %v", err)
	}
}

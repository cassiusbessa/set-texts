package handlers

import (
	"net/http"

	"github.com/cassiusbessa/db-texts/entities"
	"github.com/cassiusbessa/db-texts/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func handleServices(c *gin.Context, sec, db string) {
	var s entities.Services
	if err := c.BindJSON(&s); err != nil {
		logrus.Errorf("Error decoding Text %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := s.Validate(); err != nil {
		logrus.Errorf("Error validating Service on %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	s.Id = primitive.NewObjectID()
	if err := repositories.Push(db, sec, s); err != nil {
		logrus.Errorf("Error Setting Text in MongoDB: %v %v", db, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Service setted successfully"})
}

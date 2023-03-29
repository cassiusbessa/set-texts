package handlers

import (
	"net/http"

	"github.com/cassiusbessa/create-text/entities"
	"github.com/cassiusbessa/create-text/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func handleAddress(c *gin.Context, sec, db string) {
	var a entities.Address
	if err := c.BindJSON(&a); err != nil {
		logrus.Errorf("Error decoding Text %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := a.Validate(); err != nil {
		logrus.Errorf("Error validating Address on %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := repositories.Set(db, sec, a); err != nil {
		logrus.Errorf("Error Setting Text in MongoDB: %v %v", db, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Address setted successfully"})
}

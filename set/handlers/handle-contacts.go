package handlers

import (
	"net/http"

	"github.com/cassiusbessa/db-texts/entities"
	"github.com/cassiusbessa/db-texts/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func handleContacts(c *gin.Context, sec, db string) {
	var cont entities.Contacts
	if err := c.BindJSON(&cont); err != nil {
		logrus.Errorf("Error decoding Text %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := cont.Validate(); err != nil {
		logrus.Errorf("Error validating Contacts on %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	if err := repositories.Set(db, sec, cont); err != nil {
		logrus.Errorf("Error Setting Text in MongoDB: %v %v", db, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Contacts setted successfully"})
}

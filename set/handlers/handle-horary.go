package handlers

import (
	"net/http"

	"github.com/cassiusbessa/db-texts/entities"
	"github.com/cassiusbessa/db-texts/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func handleHorary(c *gin.Context, sec, db string) {
	var h entities.Week
	if err := c.BindJSON(&h); err != nil {
		logrus.Errorf("Error decoding Text %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if len(h.Content) != 7 {
		logrus.Errorf("Error validating Days horary on %v: %v", db, "Invalid number of days")
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid number of days"})
		return
	}
	for _, d := range h.Content {
		if err := d.Validate(); err != nil {
			logrus.Errorf("Error validating Days horary on %v: %v", db, err)
			c.JSON(http.StatusBadRequest, gin.H{"error": err})
			return
		}
	}
	if err := repositories.Set(db, sec, h.Content); err != nil {
		logrus.Errorf("Error Setting Text in MongoDB: %v %v", db, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Horary setted successfully"})
}

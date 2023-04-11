package handlers

import (
	"net/http"

	"github.com/cassiusbessa/db-texts/entities"
	"github.com/cassiusbessa/db-texts/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func handleDefaultText(c *gin.Context, sec, db string) {
	var req entities.TextSection
	if err := c.BindJSON(&req); err != nil {
		logrus.Errorf("Error decoding Text %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := repositories.Set(db, sec, req.Content); err != nil {
		logrus.Errorf("Error Setting Text in MongoDB: %v %v", db, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Text setted successfully"})
}

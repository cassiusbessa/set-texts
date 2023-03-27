package handlers

import (
	"net/http"

	"github.com/cassiusbessa/create-text/entities"
	"github.com/cassiusbessa/create-text/logs"
	"github.com/cassiusbessa/create-text/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Set(c *gin.Context) {
	defer logs.Elapsed("Create Service")()
	var req Req
	db, sec := c.Param("company"), c.Param("section")
	validSection := false
	for _, s := range entities.DisplayedSections {
		if s == sec {
			validSection = true
			break
		}
	}
	if !validSection {
		logrus.Errorf("Invalid section %v", sec)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid section"})
		return
	}
	logrus.Warnf("Setting %v Text on %v", sec, db)
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
	c.JSON(http.StatusCreated, gin.H{"message": "Text setted successfully"})
}

package handlers

import (
	"fmt"
	"net/http"

	"github.com/cassiusbessa/create-text/entities"
	"github.com/cassiusbessa/create-text/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func responser(c *gin.Context, entity entities.BaseEntity, sec, db, id string) {
	if err := c.BindJSON(&entity); err != nil {
		logrus.Errorf("Error decoding Text %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := entity.Validate(); err != nil {
		logrus.Errorf("Error validating %v on %v: %v", sec, db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	result, err := repositories.Update(db, id, sec, entity)
	if err != nil {
		logrus.Errorf("Error updating %v on %v: %v", sec, db, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": fmt.Sprintf("Error updating %v", sec)})
		return
	}
	if !result {
		logrus.Errorf("Service not found on %v, with id: %v", db, entity.GetId().String())
		c.JSON(http.StatusNotFound, gin.H{"error": fmt.Sprintf("%v not found", sec)})
		return
	}
	logrus.Infof("%v setted successfully on %v", sec, db)
	c.JSON(http.StatusOK, gin.H{"message": fmt.Sprintf("%v setted successfully", sec)})
}

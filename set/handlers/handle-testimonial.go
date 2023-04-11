package handlers

import (
	"net/http"

	"github.com/cassiusbessa/db-texts/entities"
	"github.com/cassiusbessa/db-texts/repositories"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

func handleTestimonial(c *gin.Context, sec, db string) {
	var ts entities.Testimonial
	if err := c.BindJSON(&ts); err != nil {
		logrus.Errorf("Error decoding Text %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request payload"})
		return
	}
	if err := ts.Validate(); err != nil {
		logrus.Errorf("Error validating Testimonial on %v: %v", db, err)
		c.JSON(http.StatusBadRequest, gin.H{"error": err})
		return
	}
	ts.Id = primitive.NewObjectID()
	if err := repositories.Push(db, sec, ts); err != nil {
		logrus.Errorf("Error Setting Text in MongoDB: %v %v", db, err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	c.JSON(http.StatusOK, gin.H{"message": "Testimonial setted successfully"})
}

package handlers

import (
	"fmt"
	"net/http"
	"strings"

	"github.com/cassiusbessa/create-text/entities"
	"github.com/cassiusbessa/create-text/logs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Handle(c *gin.Context) {
	defer logs.Elapsed("Set Text")()
	db, sec, id := c.Param("company"), c.Param("section"), c.Param("id")
	logrus.Warnf("Setting %v Text on %v", sec, db)
	m := fmt.Sprintf("Section must be one of: %s", strings.Join(entities.DisplayedSections, ", "))
	entityMap := map[string]entities.BaseEntity{
		"testimonials": &entities.Testimonial{},
		"services":     &entities.Service{},
	}
	e, ok := entityMap[sec]
	if !ok {
		logrus.Errorf("Invalid section %v, on %v", sec, db)
		c.JSON(http.StatusBadRequest, gin.H{"error": m})
		return
	}
	responser(c, e, sec, db, id)
}

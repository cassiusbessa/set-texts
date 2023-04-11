package handlers

import (
	"net/http"

	"github.com/cassiusbessa/db-texts/entities"
	"github.com/cassiusbessa/db-texts/logs"
	"github.com/gin-gonic/gin"
	"github.com/sirupsen/logrus"
)

func Handle(c *gin.Context) {
	defer logs.Elapsed("Set Text")()
	db, sec := c.Param("company"), c.Param("section")
	logrus.Warnf("Setting %v Text on %v", sec, db)
	validSection := false
	m := "Section must be: "
	for _, s := range entities.DisplayedSections {
		m += s + " "
		if s == sec {
			validSection = true
			break
		}
	}
	if !validSection {
		logrus.Errorf("Invalid section %v", sec)
		c.JSON(http.StatusBadRequest, gin.H{"error": m})
		return
	}
	switch sec {
	case "horary":
		handleHorary(c, sec, db)
		return
	case "testimonials":
		handleTestimonial(c, sec, db)
		return
	case "address":
		handleAddress(c, sec, db)
		return
	case "contacts":
		handleContacts(c, sec, db)
		return
	case "services":
		handleServices(c, sec, db)
		return
	default:
		handleDefaultText(c, sec, db)
		return
	}
}

package controllers

import (
	"github.com/0xc0d/weeconn/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
)

type RedirectController struct{}

func (s *RedirectController) Redirect(c *gin.Context) {
	id := c.Param("id")
	value, err := models.GetUrlByID(id)

	switch err {
	case nil:
		models.Seen(id)
	case models.Nil:
		c.JSON(http.StatusNotFound, gin.H{"error": "link not found"})
		return
	default:
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.Redirect(http.StatusMovedPermanently, value)
}

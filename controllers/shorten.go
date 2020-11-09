package controllers

import (
	"github.com/0xc0d/weeconn/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"net/url"
)

type ShortenController struct{}

func (s *ShortenController) Short(c *gin.Context) {
	value := c.PostForm("url")
	URL, err := url.ParseRequestURI(value)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "bad url"})
		return
	}
	id, err := models.AddUrl(URL.String())
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"long_url": URL.String(), "id": id})
}

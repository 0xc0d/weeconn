package controllers

import (
	"github.com/0xc0d/weeconn/models"
	"github.com/gin-gonic/gin"
	"log"
	"net/http"
	"strconv"
)

type StatisticsController struct{}

func (s *StatisticsController) Stat(c *gin.Context) {
	id := c.PostForm("id")

	period, _ := strconv.Atoi(c.Query("period"))
	seen, err := models.Stat(id, period)
	if err != nil {
		log.Println(err)
		c.JSON(http.StatusInternalServerError, gin.H{"error": "internal server error"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"id": id, "count": seen, "period": period})
}

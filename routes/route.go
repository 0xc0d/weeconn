package routes

import (
	"github.com/0xc0d/weeconn/controllers"
	"github.com/gin-gonic/gin"
)

func NewRouter() *gin.Engine {
	router := gin.New()

	redirect := new(controllers.RedirectController)
	router.GET("/:id", redirect.Redirect)

	linkGroup := router.Group("link")
	{
		shorten := new(controllers.ShortenController)
		linkGroup.POST("/shorten", shorten.Short)

	}

	statusGroup := router.Group("status")
	{
		health := new(controllers.HealthController)
		statusGroup.HEAD("/health", health.Status)

		seen := new(controllers.StatisticsController)
		statusGroup.POST("/id", seen.Stat)
	}

	return router
}

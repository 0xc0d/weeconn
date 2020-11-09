package main

import (
	"github.com/0xc0d/weeconn/routes"
	"github.com/gin-gonic/gin"
)

func main() {
	router := routes.NewRouter()
	router.Use(gin.Logger())
	router.Use(gin.Recovery())
	router.Run(":8080")
}

package main

import (
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()
	router.GET("/ping", func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	})

	staticGroup(router.Group("/static"))
	router.GET("/version", latestVersion)

	router.Run("0.0.0.0:3301")
}

func staticGroup(router *gin.RouterGroup) {
	router.Static("/update", "./update")
}

func latestVersion(c *gin.Context) {
	c.String(200, "1.0.0")
}

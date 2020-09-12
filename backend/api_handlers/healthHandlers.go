package api_handlers

import "github.com/gin-gonic/gin"

func healthHandlers() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "health",
		})
	}
}

func pingHandler() func(c *gin.Context) {
	return func(c *gin.Context) {
		c.JSON(200, gin.H{
			"message": "pong",
		})
	}
}

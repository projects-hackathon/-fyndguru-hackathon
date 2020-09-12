package api_handlers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func InitApis() *gin.Engine {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	r := gin.Default()
	r.GET("/ping", pingHandler())
	r.GET("/health", healthHandlers())
	return r
}

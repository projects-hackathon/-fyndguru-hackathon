package api_handlers

import (
	"fyndguru-hackathon/backend/api_handlers/routers"
	"github.com/gin-gonic/gin"
	"log"
)

func InitApis() *gin.Engine {
	// stdout and stderr are sent to AWS CloudWatch Logs
	log.Printf("Gin cold start")
	router := gin.Default()

	//	user login
	userRouter := router.Group("/user")
	userRouter.POST("/login", routers.Login)
	userRouter.POST("/register", routers.UserRegister)

	//employer login
	employerRouter := router.Group("/employer")
	employerRouter.POST("/login", routers.EmployerLogin)
	employerRouter.POST("/register", routers.EmployerRegister)

	router.GET("/ping", pingHandler())
	router.GET("/health", healthHandlers())
	return router
}

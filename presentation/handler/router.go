package handler

import "github.com/gin-gonic/gin"

func InitRouting(taskHandler TaskHandler) {
	router := gin.Default()

	router.POST("/task", taskHandler.Post)
	router.GET("/task/:id", taskHandler.Get)
	router.PUT("/task/:id", taskHandler.Put)
	router.DELETE("/task/:id", taskHandler.Delete)

	router.Run(":8080")
}

package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rauulssanchezz/go-gin-api/internal/handler"
)

func SetUpRoutes(router *gin.Engine, taskHandler *handler.TaskHandlerStruct) {
	api := router.Group("/api")

	{
		api.POST("/tasks", taskHandler.Create)
		api.PUT("/tasks/:id", taskHandler.Update)
		api.GET("/tasks", taskHandler.GetAll)
		api.GET("/tasks/:id", taskHandler.GetById)
		api.DELETE("/tasks/:id", taskHandler.Delete)
	}
}

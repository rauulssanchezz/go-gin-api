package router

import (
	"github.com/gin-gonic/gin"
	"github.com/rauulssanchezz/go-gin-api/internal/task"
	"github.com/rauulssanchezz/go-gin-api/internal/user"
)

func SetUpRoutes(router *gin.Engine, taskHandler task.TaskHandlerStruct, userHandler user.UserHandlerStruct) {
	api := router.Group("/api")

	{
		api.POST("/tasks", taskHandler.Create)
		api.PUT("/tasks/:id", taskHandler.Update)
		api.GET("/tasks", taskHandler.GetAll)
		api.GET("/tasks/:id", taskHandler.GetById)
		api.DELETE("/tasks/:id", taskHandler.Delete)

		api.POST("/users", userHandler.Create)
		api.PUT("/users/:id", userHandler.Update)
		api.GET("/users/:id", userHandler.GetById)
		api.GET("/users/:email", userHandler.GetByEmail)
		api.POST("/users", userHandler.Login)
		api.DELETE("/users/:id", userHandler.Delete)
	}
}

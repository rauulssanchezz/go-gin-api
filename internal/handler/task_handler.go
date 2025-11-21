package handler

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/rauulssanchezz/go-gin-api/internal/model"
	"github.com/rauulssanchezz/go-gin-api/internal/service"
)

type TaskHandler interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	Delete(context *gin.Context)
	GetAll()
	GetById(context *gin.Context)
}

type TaskHandlerStruct struct {
	Service *service.TaskServiceStruct
}

func NewTaskHandler(service *service.TaskServiceStruct) *TaskHandlerStruct {
	return &TaskHandlerStruct{
		Service: service,
	}
}

func (handler *TaskHandlerStruct) Create(context *gin.Context) {
	var task model.Task

	err := context.ShouldBindJSON(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.Service.Create(&task)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}

func (handler *TaskHandlerStruct) Update(context *gin.Context) {
	var task model.Task

	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id required"})
		return
	}

	err := context.ShouldBindJSON(&task)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.Service.Update(id, &task)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}

func (handler *TaskHandlerStruct) GetById(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id required"})
		return
	}

	task, err := handler.Service.GetById(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, task)
}

func (handler *TaskHandlerStruct) Delete(context *gin.Context) {
	id := context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "id required"})
		return
	}

	err := handler.Service.Delete(id)

	if err != nil {
		context.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}

func (handler *TaskHandlerStruct) GetAll(context *gin.Context) {
	tasks, err := handler.Service.GetAll()

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
	}

	context.JSON(http.StatusOK, tasks)
}

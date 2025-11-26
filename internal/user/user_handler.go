package user

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	GetById(context *gin.Context)
	Delete(context *gin.Context)
	Login(context *gin.Context)
	GetByEmail(context *gin.Context)
}

type UserHandlerStruct struct {
	Service UserServiceStruct
}

func NewUserHandler(service UserServiceStruct) UserHandlerStruct {
	return UserHandlerStruct{
		Service: service,
	}
}

func (handler *UserHandlerStruct) Create(context *gin.Context) {
	var user User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.Service.Create(user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}

func (handler *UserHandlerStruct) Update(context *gin.Context) {
	var id string = context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "missing required params"})
		return
	}

	var user User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	err = handler.Service.Update(id, user)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}

func (handler *UserHandlerStruct) GetById(context *gin.Context) {
	var id string = context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "missing required params"})
		return
	}

	var user User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userResponse UserResponse
	userResponse, err = handler.Service.GetById(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, userResponse)
}

func (handler *UserHandlerStruct) Delete(context *gin.Context) {
	var id string = context.Param("id")

	if id == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "missing required params"})
		return
	}

	err := handler.Service.Delete(id)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.Status(http.StatusNoContent)
}

func (handler *UserHandlerStruct) Login(context *gin.Context) {
	var credentials Credentials

	err := context.ShouldBindJSON(&credentials)
	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	if credentials.Email == "" || credentials.Password == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "missing required params"})
		return
	}

	var userResponse UserResponseLogin
	userResponse, err = handler.Service.Login(credentials.Email, credentials.Password)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.SetCookie("token", userResponse.Token, 0, "/", "", false, true)
	context.JSON(http.StatusOK, userResponse)
}

func (handler *UserHandlerStruct) GetByEmail(context *gin.Context) {
	var email string = context.Param("email")

	if email == "" {
		context.JSON(http.StatusBadRequest, gin.H{"error": "missing required params"})
		return
	}

	var user User
	err := context.ShouldBindJSON(&user)

	if err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"error": err.Error()})
		return
	}

	var userResponse UserResponse
	userResponse, err = handler.Service.GetByEmail(email)

	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}

	context.JSON(http.StatusOK, userResponse)
}

package user

import (
	"github.com/gin-gonic/gin"
)

type UserHandler interface {
	Create(context *gin.Context)
	Update(context *gin.Context)
	GetById(context *gin.Context)
	Delete(context *gin.Context)
}

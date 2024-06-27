package controllers

import (
	"net/http"
	"github.com/gin-gonic/gin"
	"github.com/xams-backend/services/auth-service/src/internal/users/usecase"
)

type (
	UserController struct {
		Usecase usecase.IUserUsecase
	}
)

func (uc *UserController) GetUser(context *gin.Context) {
	userId := context.Param("id")
	if userId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid parameter"})
		return
	}

	userResponse, err := uc.Usecase.GetUser(userId)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, userResponse)
}
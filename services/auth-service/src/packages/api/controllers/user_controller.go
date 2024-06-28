package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xams-backend/services/auth-service/src/internal/models"
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

func (uc *UserController) UpdateUser(context *gin.Context) {
	var request models.User

	userId := context.Param("id")
	if userId == "" {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid parameter"})
		return
	}

	if err := uc.Usecase.IsUserAlreadyExists(userId); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": "invalid parameter"})
		return
	}

	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	if err := uc.Usecase.UpdateUser(&request); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	context.JSON(http.StatusOK, gin.H{"message": "update user successfully"})
}
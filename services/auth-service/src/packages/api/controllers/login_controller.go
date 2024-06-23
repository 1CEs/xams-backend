package controllers

import (
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/xams-backend/services/auth-service/src/internal/models"
	"github.com/xams-backend/services/auth-service/src/packages/auth"
)

type (
	LoginController struct {
		Auth auth.IAuthentication
	}
)

func (lc *LoginController) Login(context *gin.Context) {
	var request models.LoginRequest
	if err := context.ShouldBind(&request); err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}
	userResponse, err := lc.Auth.Login(request.Email, request.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err})
		return
	}

	context.JSON(http.StatusOK, userResponse)
}
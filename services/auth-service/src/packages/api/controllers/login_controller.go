package controllers

import (
	"encoding/json"
	"log"
	"net/http"
	"os"

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
	log.Println("in controller")
	var request models.LoginRequest
	if err := context.ShouldBind(&request); err != nil {
		log.Println("bind error")
		context.JSON(http.StatusBadRequest, gin.H{"message": err.Error()})
		return
	}
	
	userJson, _ := json.MarshalIndent(request, "", "		")
	log.Println(string(userJson))

	userResponse, token, err := lc.Auth.Login(request.Email, request.Password)
	if err != nil {
		context.JSON(http.StatusInternalServerError, gin.H{"message": err.Error()})
		return
	}

	cookieName := os.Getenv("COOKIE_JWT_TOKEN")

	context.SetSameSite(http.SameSiteLaxMode)
	context.SetCookie(cookieName, token, 3600 * 24 * 30, "", "", false, true)

	context.JSON(http.StatusOK, userResponse)
}
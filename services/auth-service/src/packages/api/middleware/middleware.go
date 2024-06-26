package middleware

import "github.com/gin-gonic/gin"

type Middleware struct{}

func (m *Middleware) IsAuthorized(context *gin.Context) {
	
}
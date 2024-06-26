package models

import "github.com/golang-jwt/jwt/v4"

type (
	LogInResponse struct {
		UserID    string `json:"user_id"`
		Email     string `json:"email"`
		Prename   string `json:"prename"`
		FirstName string `json:"first_name"`
		LastName  string `json:"last_name"`
		BranchID  uint   `json:"branch_id"`
		Role      Role   `json:"role"`
	}

	UserClaims struct {
		UserID    string 
		Role      Role
		jwt.RegisteredClaims
	}

	LoginRequest struct {
		Email     string `json:"email" binding:"required"`
		Password  string `json:"password" binding:"required"`
	}

)

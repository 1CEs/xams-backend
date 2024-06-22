package models

type (
	UserResponse struct {
		UserID    string   `json:"user_id"`
		Email     string   `json:"email"`
		Prename   string   `json:"prename"`
		FirstName string   `json:"first_name"`
		LastName  string   `json:"last_name"`
		BranchID  uint     `json:"branch_id"`
		Role      Role     `json:"role"`
		//Token     JWTToken `json:"token"`
	}

	JWTToken struct {
		AccessToken  string `json:"access_token"`
		RefreshToken string `json:"refresh_token"`
		ExpiresIn    int64  `json:"expires_in"`
	}
)

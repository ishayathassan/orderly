package models

type RegisterRequest struct {
	Username string `json:"username" binding:"required" example:"johndoe"`
	Password string `json:"password" binding:"required" example:"Secret123!"`
}

type RegisterResponse struct {
	ID       string `json:"id" example:"1a2b3c4d"`
	Username string `json:"username" example:"johndoe"`
	Role     string `json:"role" example:"user"`
}
package models

type ErrorResponse struct {
	Code    string `json:"code,omitempty" example:"INVALID_REQUEST"`
	Message string `json:"message" example:"Invalid username or password"`
}
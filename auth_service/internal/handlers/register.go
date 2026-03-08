package handlers

import (
	"errors"
	"net/http"
	"orderly/auth-service/internal/models"
	"orderly/auth-service/internal/services"
	"orderly/auth-service/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Register user
// @Description Create a new user
// @Tags auth
// @Accept json
// @Produce json
// @Param register body models.RegisterRequest true "Register Request"
// @Success 201 {object} models.RegisterResponse "Successful Registration Example"
// @Failure 400 {object} models.ErrorResponse "Invalid Request Example"
// @Failure 409 {object} models.ErrorResponse "Username Already Exists Example"
// @Router /register [post]
func Register(c *gin.Context) {
	var req models.RegisterRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	user, err := services.Register(req.Username, req.Password)
	if err != nil {
		if errors.Is(err, utils.ErrUsernameExists) {
			utils.RespondError(c, http.StatusConflict, "USERNAME_EXISTS", "Username already exists")
			return
		}
		utils.RespondError(c, http.StatusInternalServerError, "INTERNAL_ERROR", "Could not create user")
		return
	}

	c.JSON(http.StatusCreated, models.RegisterResponse{
		ID:       user.ID.String(),
		Username: user.Username,
		Role:     user.Role,
	})
}
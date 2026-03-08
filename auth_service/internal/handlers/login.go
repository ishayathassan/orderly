package handlers

import (
	"net/http"
	"orderly/auth-service/internal/models"
	"orderly/auth-service/internal/services"
	"orderly/auth-service/utils"

	"github.com/gin-gonic/gin"
)

// @Summary Login user
// @Description Login with username and password
// @Tags auth
// @Accept json
// @Produce json
// @Param login body models.LoginRequest true "Login Request"
// @Success 200 {object} models.LoginResponse "Successful Login Example"
// @Failure 400 {object} models.ErrorResponse "Invalid Request Example"
// @Failure 401 {object} models.ErrorResponse "Unauthorized Example"
// @Router /login [post]
func Login(c *gin.Context) {
	var req models.LoginRequest
	if err := c.ShouldBindJSON(&req); err != nil {
		utils.RespondError(c, http.StatusBadRequest, "INVALID_REQUEST", err.Error())
		return
	}

	token, err := services.Login(req.Username, req.Password)
	if err != nil {
		utils.RespondError(c, http.StatusUnauthorized, "INVALID_CREDENTIALS", "Invalid username or password")
		return
	}

	c.JSON(http.StatusOK, models.LoginResponse{Token: token})
}

package services

import (
	"orderly/auth-service/internal/auth"
	"orderly/auth-service/internal/metrics"
	"orderly/auth-service/internal/models"
	"orderly/auth-service/internal/repository"
	"orderly/auth-service/utils"
	"strings"
)


func Register(username, password string) (*models.User, error) {

	hashedPassword, err := auth.HashPassword(password)
	if err != nil {
		return nil, err
	}

	user := models.User{
		Username: username,
		Password: hashedPassword,
		Role:     "user",
	}

	err = repository.CreateUser(&user)
	if err != nil {
		if strings.Contains(err.Error(), "duplicate") {
			return nil, utils.ErrUsernameExists
		}
		return nil, err
	}

	metrics.UsersCreatedTotal.Inc()
	return &user, nil
}




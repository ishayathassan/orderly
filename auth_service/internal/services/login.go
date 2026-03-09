package services

import (
	"errors"
	"orderly/auth-service/internal/auth"
	"orderly/auth-service/internal/metrics"
	"orderly/auth-service/internal/repository"
)

func Login(username, password string) (string, error) {

	user, err := repository.GetUserByUsername(username)
	if err != nil {
		return "", errors.New("invalid credentials")
	}

	if !auth.CheckPasswordHash(password, user.Password) {
		return "", errors.New("invalid credentials")
	}

	token, err := auth.GenerateToken(user.ID.String(), user.Role)
	if err != nil {
		return "", err
	}

	metrics.TokensIssuedTotal.Inc()
	return token, nil
}
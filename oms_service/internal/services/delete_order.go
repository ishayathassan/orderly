package services

import (
	"orderly/oms-service/internal/repositories"
	"orderly/oms-service/internal/utils"
)

func DeleteOrder(id uint) error {
	_, err := repositories.GetByID(id)
	if err != nil {
		return utils.ErrOrderNotFound
	}
	return repositories.Delete(id)
}
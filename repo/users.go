package repo

import (
	"gorm.io/gorm"

	"rbac/models"
)

type User interface {
	Create(models.User) (*models.User, error)
}

type userImpl struct {
	db *gorm.DB
}

func (repo *userImpl) Create(user models.User) (*models.User, error) {
	err := repo.db.Create(&user).Error
	if err != nil {
		return nil, err
	}

	return &user, nil
}

package persistence

import (
	"gofiber-orm/domain/model"
	"gorm.io/gorm"
)

type UserRepository struct {
	DB *gorm.DB
}

func (r *UserRepository) GetUserByCookie(cookie string) (*model.User, error) {
	var user model.User
	result := r.DB.Where("username = ?", cookie).First(&user)
	if result.Error != nil {
		return nil, result.Error
	}

	return &user, nil
}

package service

import (
	"errors"
	"gofiber-orm/domain/model"
	"gofiber-orm/domain/repository"
)

type UserService interface {
	GetUser(id uint64) (*model.User, error)
}

type userService struct {
	userRepo repository.UserRepository
}

func NewUserService(userRepo repository.UserRepository) UserService {
	return &userService{userRepo}
}

func (s *userService) GetUser(id uint64) (*model.User, error) {
	user, err := s.userRepo.FindById(id)
	if err != nil {
		return nil, errors.New("user not found")
	}
	return user, nil
}

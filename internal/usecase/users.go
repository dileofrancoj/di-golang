package usecase

import (
	"di-golang/internal/models"
	"di-golang/internal/repository"
)

type UserUseCase interface {
	GetAll() ([]models.User, error)
}

type userUseCase struct {
	repo repository.UserRepository
}

func NewUserUseCase(repo repository.UserRepository) UserUseCase {
	return &userUseCase{repo: repo}
}

func (u *userUseCase) GetAll() ([]models.User, error) {
	return u.repo.GetAll()
}

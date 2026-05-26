package usecase

import "di-golang/internal/models"


type UserUseCase interface {
	GetAll() []models.User
}

type userUseCase struct{}

func NewUserUseCase() UserUseCase {
	return &userUseCase{}
}

func (u *userUseCase) GetAll() []models.User {
	return []models.User{
		{ID: 1, Name: "User One", Email: "user1@example.com"},
		{ID: 2, Name: "User Two", Email: "user2@example.com"},
	}
}

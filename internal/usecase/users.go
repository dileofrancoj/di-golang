package usecase

import "di-golang/internal/domain"



type UserUseCase interface {
	GetAll() []domain.User
}

type userUseCase struct{}

func NewUserUseCase() UserUseCase {
	return &userUseCase{}
}

func (u *userUseCase) GetAll() []domain.User {
	return []domain.User{
		{ID: 1, Name: "User One", Email: "user1@example.com"},
		{ID: 2, Name: "User Two", Email: "user2@example.com"},
	}
}

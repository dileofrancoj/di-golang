package repository

import (
	"di-golang/internal/models"
	"encoding/json"
	"os"
)

type UserRepository interface {
	GetAll() ([]models.User, error)
}

type userRepository struct {
	filePath string
}

func NewUserRepository(filePath string) UserRepository {
	return &userRepository{filePath: filePath}
}

func (r *userRepository) GetAll() ([]models.User, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var users []models.User
	if err := json.Unmarshal(data, &users); err != nil {
		return nil, err
	}

	return users, nil
}

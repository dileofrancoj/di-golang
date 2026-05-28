package repository

import (
	"di-golang/internal/models"
	"encoding/json"
	"errors"
	"os"
)

type ProductRepository interface {
	GetAll() ([]models.Product, error)
	GetByID(id int) (*models.Product, error)
}

type productRepository struct {
	filePath string
}

func NewProductRepository(filePath string) ProductRepository {
	return &productRepository{filePath: filePath}
}

func (r *productRepository) GetAll() ([]models.Product, error) {
	data, err := os.ReadFile(r.filePath)
	if err != nil {
		return nil, err
	}

	var products []models.Product
	if err := json.Unmarshal(data, &products); err != nil {
		return nil, err
	}

	return products, nil
}

func (r *productRepository) GetByID(id int) (*models.Product, error) {
	products, err := r.GetAll()
	if err != nil {
		return nil, err
	}

	for _, p := range products {
		if p.ID == id {
			return &p, nil
		}
	}

	return nil, errors.New("product not found")
}

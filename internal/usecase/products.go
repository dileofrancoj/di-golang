package usecase

import (
	"di-golang/internal/models"
	"di-golang/internal/repository"
)

type ProductUseCase interface {
	GetAll() ([]models.Product, error)
	GetProduct(id int) (*models.Product, error)
}

type productUseCase struct {
	repo repository.ProductRepository
}

func NewProductUseCase(repo repository.ProductRepository) ProductUseCase {
	return &productUseCase{repo: repo}
}

func (p *productUseCase) GetProduct(id int) (*models.Product, error) {
	return p.repo.GetByID(id)
}

func (p *productUseCase) GetAll() ([]models.Product, error) {
	return p.repo.GetAll()
}

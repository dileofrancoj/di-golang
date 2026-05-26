package usecase

import "di-golang/internal/models"

type ProductUseCase interface {
	GetAll() []models.Product
}


type productUseCase struct{}

func NewProductUseCase() ProductUseCase {
	return &productUseCase{}
}

func (p *productUseCase) GetAll() []models.Product {
	return []models.Product{
		{ID: 1, Name: "Product A", Price: 10.5},
		{ID: 2, Name: "Product B", Price: 20.0},
	}
}

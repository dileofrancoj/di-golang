package usecase

import "di-golang/internal/domain"

type ProductUseCase interface {
	GetAll() []domain.Product
}


type productUseCase struct{}

func NewProductUseCase() ProductUseCase {
	return &productUseCase{}
}

func (p *productUseCase) GetAll() []domain.Product {
	return []domain.Product{
		{ID: 1, Name: "Product A", Price: 10.5},
		{ID: 2, Name: "Product B", Price: 20.0},
	}
}

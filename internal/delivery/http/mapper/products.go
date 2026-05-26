package mapper

import (
	"fmt"
	"di-golang/internal/delivery/http/dto"
	"di-golang/internal/domain"
)

func ToProductResponseList(products []domain.Product) []dto.ProductResponse {
	fmt.Println(products)
	response := make([]dto.ProductResponse , len(products))
	for i, p := range products {
		response[i] = dto.ProductResponse{
			ID:    p.ID,
			Name:  p.Name,
			Price: p.Price,
		}
	}
	return response
}
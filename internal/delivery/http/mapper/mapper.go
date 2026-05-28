package mapper

import (
	"di-golang/internal/delivery/http/dto"
	"di-golang/internal/models"
)

func ToProductResponseList(domainProducts []models.Product) []dto.ProductResponse {
	responses := make([]dto.ProductResponse, len(domainProducts))
	for i, p := range domainProducts {
		responses[i] = ToProductResponse(p)
	}
	return responses
}

func ToProductResponse(p models.Product) dto.ProductResponse {
	return dto.ProductResponse{
		ID:    p.ID,
		Name:  p.Name,
		Price: p.Price,
	}
}

func ToUserResponseList(domainUsers []models.User) []dto.UserResponse {
	responses := make([]dto.UserResponse, len(domainUsers))
	for i, u := range domainUsers {
		responses[i] = dto.UserResponse{
			ID:    u.ID,
			Name:  u.Name,
			Email: u.Email,
		}
	}
	return responses
}

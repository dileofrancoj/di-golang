package mapper

import (
	"di-golang/internal/delivery/http/dto"
	"di-golang/internal/models"
)


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

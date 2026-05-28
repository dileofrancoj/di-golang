package handlers

import (
	"di-golang/internal/usecase"
)

type Handler struct {
	productUC usecase.ProductUseCase
	userUC    usecase.UserUseCase
}

func NewHandler(pUC usecase.ProductUseCase, uUC usecase.UserUseCase) *Handler {
	return &Handler{
		productUC: pUC,
		userUC:    uUC,
	}
}
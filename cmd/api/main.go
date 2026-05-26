package main

import (
	"log"

	delivery "di-golang/internal/delivery/http"
	"di-golang/internal/handlers"
	uc "di-golang/internal/usecase"
)

func main() {
	// Initialize Use Cases
	productUC := uc.NewProductUseCase()
	userUC := uc.NewUserUseCase()

	h := handlers.NewHandler(productUC, userUC)

	r := delivery.SetupRouter(h)

	log.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}


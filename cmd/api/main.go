package main

import (
	"log"

	delivery "di-golang/internal/delivery/http"
	"di-golang/internal/handlers"
	"di-golang/internal/repository"
	uc "di-golang/internal/usecase"
)

func main() {
	// Initialize Repositories
	productRepo := repository.NewProductRepository("data/products.json")
	userRepo := repository.NewUserRepository("data/users.json")

	// Initialize Use Cases
	productUC := uc.NewProductUseCase(productRepo)
	userUC := uc.NewUserUseCase(userRepo)

	// Initialize Handlers
	h := handlers.NewHandler(productUC, userUC)

	// Setup Router
	r := delivery.SetupRouter(h)

	// Start Server
	log.Println("Server starting on :8080...")
	if err := r.Run(":8080"); err != nil {
		log.Fatal(err)
	}
}

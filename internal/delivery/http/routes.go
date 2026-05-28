package http

import (
	"di-golang/internal/handlers"
	"github.com/gin-gonic/gin"
)

func SetupRouter(h *handlers.Handler) *gin.Engine {
	r := gin.Default()

	r.GET("/ping", h.Ping)
	r.GET("/products", h.GetProducts)
	r.GET("/products/:id", h.GetProduct)
	r.GET("/users", h.GetUsers)

	return r
}

package handlers

import (
	"net/http"

	"di-golang/internal/delivery/http/mapper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProducts(c *gin.Context) {
	products := h.productUC.GetAll()
	response := mapper.ToProductResponseList(products)
	c.JSON(http.StatusOK, response)
}

package handlers

import (
	"net/http"
	"strconv"

	"di-golang/internal/delivery/http/mapper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetProducts(c *gin.Context) {
	products, err := h.productUC.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := mapper.ToProductResponseList(products)
	c.JSON(http.StatusOK, response)
}

func (h *Handler) GetProduct(c *gin.Context) {
	idParam := c.Param("id")
	id, err := strconv.Atoi(idParam)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "invalid product id"})
		return
	}

	product, err := h.productUC.GetProduct(id)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": err.Error()})
		return
	}

	response := mapper.ToProductResponse(*product)
	c.JSON(http.StatusOK, response)
}

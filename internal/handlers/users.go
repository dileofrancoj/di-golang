package handlers

import (
	"net/http"

	"di-golang/internal/delivery/http/mapper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {
	users, err := h.userUC.GetAll()
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
		return
	}
	response := mapper.ToUserResponseList(users)
	c.JSON(http.StatusOK, response)
}

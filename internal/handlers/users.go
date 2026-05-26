package handlers

import (
	"net/http"

	"di-golang/internal/delivery/http/mapper"

	"github.com/gin-gonic/gin"
)

func (h *Handler) GetUsers(c *gin.Context) {
	users := h.userUC.GetAll()
	response := mapper.ToUserResponseList(users)
	c.JSON(http.StatusOK, response)
}

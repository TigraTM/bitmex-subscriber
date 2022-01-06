package web

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func (h *Handler) Create(c *gin.Context) {
	var req CreateRequest
	if err := c.BindJSON(&req); err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	user, err := h.app.Create(c, req.Name)
	if err != nil {
		c.AbortWithStatusJSON(http.StatusBadRequest, err)
		return
	}

	c.JSON(http.StatusCreated, user)
}

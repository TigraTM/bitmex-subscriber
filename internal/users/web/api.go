package web

import (
	"bitmex-subscriber/internal/users/app"
	"context"
	"github.com/gin-gonic/gin"
)

type (
	application interface {
		Create(ctx context.Context, name string) (app.User, error)
	}

	Handler struct {
		app application
	}
)

func New(app application) *Handler {
	return &Handler{
		app: app,
	}
}

func (h *Handler) Init(group *gin.RouterGroup) {
	users := group.Group("/users")

	{
		users.POST("/", h.Create)
	}
}

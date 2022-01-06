package web

import (
	"github.com/google/uuid"
)

type CreateRequest struct {
	Name string `json:"name"`
}

type CreateResponse struct {
	ID   uuid.UUID `json:"id"`
	Name string    `json:"name"`
}

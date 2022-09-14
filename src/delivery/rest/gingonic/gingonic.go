package gingonic

import (
	"github.com/dalikewara/ayapingping-go/v2/src/service"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	service *service.Service
	client  *gin.Engine
	config  *Config
}

// New generates new gin-gonic handler.
func New(param NewParam) Handler {
	panic("implement me")
}

// Serve serves gin-gonic application.
func (h *Handler) Serve() {
	panic("implement me")
}

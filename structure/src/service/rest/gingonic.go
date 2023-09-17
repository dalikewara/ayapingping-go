package rest

import (
	"github.com/gin-gonic/gin"
)

type ginGonicHandler struct {
	client *gin.Engine
}

func (g *ginGonicHandler) Serve() {
	panic("implement me")
}

// New

func NewGinGonic(client *gin.Engine) Contract {
	return &ginGonicHandler{
		client: client,
	}
}

package ping

import (
	"github.com/gin-gonic/gin"
)

type GinGonic struct{}

func (g *GinGonic) String(c *gin.Context) {
	panic("implement me")
}

// New

func NewGinGonic() *GinGonic {
	return &GinGonic{}
}

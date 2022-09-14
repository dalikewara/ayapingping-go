package gingonic

import (
	"github.com/dalikewara/ayapingping-go/v2/src/service"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Port string
}

type NewParam struct {
	Service *service.Service
	Client  *gin.Engine
	Config  *Config
}

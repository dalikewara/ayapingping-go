package rest

import (
	"github.com/dalikewara/ayapingping-go/src/service"
	"github.com/gin-gonic/gin"
)

type Config struct {
	Env             string
	Port            string
	OkCode          string
	OkMessage       string
	OkHttpStatus    int
	NotOkHttpStatus int
}

type NewGinParam struct {
	Service *service.Service
	Client  *gin.Engine
	Config  *Config
}

type Response struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type APIUserV1GetAllData struct {
	Id       int    `json:"id"`
	Username string `json:"username"`
}

type APIRoleV1GetByUserIDData struct {
	Id   int    `json:"id"`
	Role string `json:"role"`
}

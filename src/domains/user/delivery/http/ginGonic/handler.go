package ginGonic

import (
	"github.com/dalikewara/ayapingping-go/src/domains/user"
	"github.com/gin-gonic/gin"
)

type Handler struct {
	router  *gin.Engine
	useCase user.UseCaseInterface
}

type NewHandlerParam struct {
	Router  *gin.Engine
	UseCase user.UseCaseInterface
}

// NewHandler generates new http Handler.
func NewHandler(param NewHandlerParam) {
	httpHandler := &Handler{
		router:  param.Router,
		useCase: param.UseCase,
	}
	httpHandler.Routes()
}

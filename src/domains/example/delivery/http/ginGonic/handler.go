package ginGonic

import (
	"github.com/dalikewara/ayapingping-go/src/domains/example"
	"github.com/gin-gonic/gin"
	"net/http"
)

type Handler struct {
	router  *gin.Engine
	useCase example.UseCaseInterface
}

type NewHandlerParam struct {
	Router  *gin.Engine
	UseCase example.UseCaseInterface
}

// NewHandler generates new http Handler.
func NewHandler(param NewHandlerParam) {
	httpHandler := &Handler{
		router:  param.Router,
		useCase: param.UseCase,
	}
	httpHandler.Routes()
}

// Routes handles API routes.
func (e *Handler) Routes() {
	e.router.GET("/test", e.Test)
}

// Test serves example API for method GetAndChangeName of example.UseCaseInterface.
func (e *Handler) Test(g *gin.Context) {
	res := e.useCase.GetAndChangeName(example.UseCaseGetAndChangeNameParam{
		Id:   1,
		Name: "Smith",
	})
	if res.Error != nil {
		g.JSON(http.StatusInternalServerError, example.HttpResponse{
			Code:    "-1",
			Message: res.Error.Error(),
			Data:    nil,
		})
	} else {
		g.JSON(http.StatusOK, example.HttpResponse{
			Code:    "00",
			Message: "Ok",
			Data: &example.HttpTestResponseData{
				Id:        res.Example.Id,
				Name:      res.Example.Name,
				IsChanged: res.IsChanged,
			},
		})
	}
}

package apis

import (
	"github.com/dalikewara/ayapingping-go/src/domains/example"
	"github.com/gin-gonic/gin"
	"net/http"
)

// APIExample struct.
type APIExample struct {
	router *gin.Engine
	exampleUseCase example.UseCaseInterface
}

// NewAPIExampleParam parameter.
type NewAPIExampleParam struct {
	Router *gin.Engine
	ExampleUseCase example.UseCaseInterface
}

// NewAPIExample generates new NewAPIExample handler.
func NewAPIExample(param NewAPIExampleParam) {
	apiExample := &APIExample{
		router: param.Router,
		exampleUseCase: param.ExampleUseCase,
	}
	apiExample.router.GET("/test", apiExample.Test)
}

// Test serves example api for getAndChangeName use case (/get).
func (a *APIExample) Test(g *gin.Context) {
	res := a.exampleUseCase.GetAndChangeName(example.UseCaseGetAndChangeNameParam{
		Id: 1,
		Name: "Smith",
	})
	if res.Error != nil {
		g.JSON(http.StatusOK, example.APITestResponse{
			Code: "01",
			Message: "Err",
			Data: nil,
		})
		return
	}
	g.JSON(http.StatusOK, example.APITestResponse{
		Code: "00",
		Message: "Ok",
		Data: &example.APITestResponseData{
			Id: res.Example.Id,
			Name: res.Example.Name,
			IsChanged: res.IsChanged,
		},
	})
}

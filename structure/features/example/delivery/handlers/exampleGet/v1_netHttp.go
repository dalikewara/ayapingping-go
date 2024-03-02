package exampleGet

import (
	"encoding/json"
	"github.com/dalikewara/ayapingping-go/v4/structure/domain"
	"github.com/dalikewara/ayapingping-go/v4/structure/features/example/commons/presenterJSON"
	"github.com/dalikewara/ayapingping-go/v4/structure/features/example/delivery/middlewares/checkMethod"
	"net/http"
)

type v1NetHttp struct {
	serverMux  *http.ServeMux
	getExample domain.GetExampleUseCase
}

func NewV1NetHttp(serverMux *http.ServeMux, getExample domain.GetExampleUseCase) domain.ExampleDelivery {
	return &v1NetHttp{
		serverMux:  serverMux,
		getExample: getExample,
	}
}

func (v *v1NetHttp) RegisterHandler(method string, endpoint string) {
	v.serverMux.Handle(endpoint, checkMethod.NetHttp(method, http.HandlerFunc(v.getQueryParam)))
}

func (v *v1NetHttp) getQueryParam(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	dto, _ := v.getExample.ExecCtx(ctx, 1)
	httpStatus, result := presenterJSON.OK("00", "", dto)
	resultBytes, _ := json.Marshal(result)

	w.WriteHeader(httpStatus)

	_, _ = w.Write(resultBytes)

	return
}

package exampleGet

import (
	"encoding/json"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/service/rest/presenter"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/usecase/getExample"
	"net/http"
)

type V1HttpServer struct {
	getExample getExample.Contract
}

// JSON handles json response
func (v *V1HttpServer) JSON(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload V1Payload

	_ = presenter.UrlQueryBindPayload(r.URL.Query(), &payload)

	exampleDTO, _ := v.getExample.ExecCtx(ctx, payload.ID)
	httpStatus, result := presenter.JSONSuccess(http.StatusOK, exampleDTO)
	resultBytes, _ := json.Marshal(result)

	w.WriteHeader(httpStatus)

	_, _ = w.Write(resultBytes)
	return
}

// NewV1HttpServer generates new v1 http server
func NewV1HttpServer(getExample getExample.Contract) *V1HttpServer {
	return &V1HttpServer{
		getExample: getExample,
	}
}

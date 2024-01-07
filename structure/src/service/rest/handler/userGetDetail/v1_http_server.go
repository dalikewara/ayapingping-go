package userGetDetail

import (
	"encoding/json"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/service/rest/presenter"
	"github.com/dalikewara/ayapingping-go/v3/structure/src/usecase/getUserDetail"
	"net/http"
)

type V1HttpServer struct {
	getUserDetail getUserDetail.Contract
}

// JSON handles JSON response
func (v *V1HttpServer) JSON(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()

	var payload V1Payload

	_ = presenter.UrlQueryBindPayload(r.URL.Query(), &payload)

	userDTO, _ := v.getUserDetail.ExecCtx(ctx, payload.UserID)
	httpStatus, result := presenter.JSONSuccess(200, userDTO)
	resultBytes, _ := json.Marshal(result)

	w.WriteHeader(httpStatus)

	_, _ = w.Write(resultBytes)
	return
}

// NewV1HttpServer creates new v1 http server
func NewV1HttpServer(getUserDetail getUserDetail.Contract) *V1HttpServer {
	return &V1HttpServer{
		getUserDetail: getUserDetail,
	}
}

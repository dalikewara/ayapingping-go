package example

import (
	"encoding/json"
	"github.com/dalikewara/ayapingping-go/v4/_base_structure/domain"
	"net/http"
	"strings"
)

type httpServiceNetHttp struct {
	client         *http.ServeMux
	exampleUseCase domain.ExampleUseCase
}

func NewHttpServiceNetHttp(client *http.ServeMux, exampleUseCase domain.ExampleUseCase) domain.ExampleHttpService {
	return &httpServiceNetHttp{
		client:         client,
		exampleUseCase: exampleUseCase,
	}
}

func (h *httpServiceNetHttp) ExampleDetail(method string, endpoint string) {
	h.client.HandleFunc(endpoint, func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()

		if r.Method != strings.ToUpper(method) {
			resultBytes, _ := json.Marshal(domain.NewResponseJSONError(ErrInvalidRequestMethod))

			w.WriteHeader(400)

			_, _ = w.Write(resultBytes)

			return
		}

		result, err := h.exampleUseCase.GetDetailCtx(ctx, 1)
		if err != nil {
			resultBytes, _ := json.Marshal(domain.NewResponseJSONError(err))

			w.WriteHeader(500)

			_, _ = w.Write(resultBytes)

			return
		}

		resultBytes, _ := json.Marshal(domain.NewResponseJSONSuccess(result))

		w.WriteHeader(200)

		_, _ = w.Write(resultBytes)

		return
	})
}

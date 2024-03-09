package checkMethod

import (
	"encoding/json"
	"github.com/dalikewara/ayapingping-go/v4/_baseStructure/features/example/commons/errMsg"
	"github.com/dalikewara/ayapingping-go/v4/_baseStructure/features/example/commons/presenterJSON"
	"net/http"
	"strings"
)

func NetHttp(method string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		if strings.ToUpper(method) != strings.ToUpper(r.Method) {
			httpStatus, result := presenterJSON.NotOK("00", "", []error{errMsg.InvalidRequestMethod})
			resultBytes, _ := json.Marshal(result)

			w.WriteHeader(httpStatus)

			_, _ = w.Write(resultBytes)
			return
		}

		next.ServeHTTP(w, r)
	})
}

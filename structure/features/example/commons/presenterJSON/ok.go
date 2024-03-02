package presenterJSON

import (
	"github.com/dalikewara/ayapingping-go/v4/structure/domain"
	"net/http"
)

func OK(code string, message string, data interface{}) (int, *domain.ExamplePresenterJSON) {
	return http.StatusOK, &domain.ExamplePresenterJSON{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

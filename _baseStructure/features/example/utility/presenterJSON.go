package utility

import (
	"github.com/dalikewara/ayapingping-go/v4/_baseStructure/domain"
	"net/http"
)

func PresenterJSONOK(code string, message string, data interface{}) (int, *domain.ExamplePresenterJSON) {
	return http.StatusOK, &domain.ExamplePresenterJSON{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func PresenterJSONNotOK(code string, message string, errs []error) (int, *domain.ExamplePresenterJSON) {
	var errText []string

	for _, err := range errs {
		if err == nil {
			continue
		}

		errText = append(errText, err.Error())
	}

	return http.StatusOK, &domain.ExamplePresenterJSON{
		Code:    code,
		Message: message,
		Errors:  errText,
	}
}

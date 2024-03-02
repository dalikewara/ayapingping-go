package presenterJSON

import (
	"github.com/dalikewara/ayapingping-go/v4/structure/domain"
	"net/http"
)

func NotOK(code string, message string, errs []error) (int, *domain.ExamplePresenterJSON) {
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

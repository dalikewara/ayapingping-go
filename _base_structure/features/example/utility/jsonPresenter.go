package utility

import (
	"github.com/dalikewara/ayapingping-go/v4/_base_structure/domain"
	"net/http"
)

func JSONPresenterOK(code string, message string, data interface{}) (int, *domain.ExampleJSONPresenter) {
	return http.StatusOK, &domain.ExampleJSONPresenter{
		Code:    code,
		Message: message,
		Data:    data,
	}
}

func JSONPresenterNotOK(code string, message string, errs []error) (int, *domain.ExampleJSONPresenter) {
	var errText []string

	for _, err := range errs {
		if err == nil {
			continue
		}

		errText = append(errText, err.Error())
	}

	return http.StatusOK, &domain.ExampleJSONPresenter{
		Code:    code,
		Message: message,
		Errors:  errText,
	}
}

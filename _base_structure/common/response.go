package common

type ResponseJSON struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
	Errors  []string    `json:"errors"`
}

func NewResponseJSONSuccess(data interface{}) *ResponseJSON {
	return &ResponseJSON{
		Status:  true,
		Message: "ok",
		Data:    data,
	}
}

func NewResponseJSONError(err error) *ResponseJSON {
	return &ResponseJSON{
		Status:  false,
		Message: "error",
		Errors: []string{
			err.Error(),
		},
	}
}

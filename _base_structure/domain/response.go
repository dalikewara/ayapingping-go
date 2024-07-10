package domain

type ResponseJSON struct {
	Status  bool        `json:"status"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

func NewResponseJSONSuccess(data interface{}) *ResponseJSON {
	return &ResponseJSON{
		Status:  true,
		Message: "success",
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

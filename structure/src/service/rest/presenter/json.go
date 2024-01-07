package presenter

type JSON struct {
	Code    string      `json:"code"`
	Message string      `json:"message"`
	Data    interface{} `json:"data,omitempty"`
	Errors  []string    `json:"errors,omitempty"`
}

// JSONSuccess builds JSON success object
func JSONSuccess(httpStatus int, data interface{}) (int, *JSON) {
	return httpStatus, &JSON{
		Code:    "00",
		Message: "ok",
		Data:    data,
	}
}

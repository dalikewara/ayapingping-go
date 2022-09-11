package errs

type Errs interface {
	GetCode() string
	GetMessage() string
	GetHttpStatus() int
}

type errs struct {
	Code       string `json:"code"`
	Message    string `json:"message"`
	HttpStatus int    `json:"http_status"`
}

// New generates new errs that implements Errs.
func New(code, message string) Errs {
	return &errs{
		Code:    code,
		Message: message,
	}
}

// NewWithHttpStatus generates new errs with http status that implements Errs.
func NewWithHttpStatus(code, message string, httpStatus int) Errs {
	return &errs{
		Code:       code,
		Message:    message,
		HttpStatus: httpStatus,
	}
}

// GetCode gets error code.
func (e *errs) GetCode() string {
	return e.Code
}

// GetMessage gets error message.
func (e *errs) GetMessage() string {
	return e.Message
}

// GetHttpStatus gets error http status.
func (e *errs) GetHttpStatus() int {
	return e.HttpStatus
}

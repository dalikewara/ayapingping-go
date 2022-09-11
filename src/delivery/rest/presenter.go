package rest

// Ok generates success response payload.
func Ok(config *Config, data interface{}) (int, Response) {
	return config.OkHttpStatus, Response{
		Code:    config.OkCode,
		Message: config.OkMessage,
		Data:    data,
	}
}

// NotOk generates error response payload.
func NotOk(config *Config, httpStatus int, code, message string) (int, Response) {
	if httpStatus != 0 {
		return httpStatus, Response{
			Code:    code,
			Message: message,
		}
	} else {
		return config.NotOkHttpStatus, Response{
			Code:    code,
			Message: message,
		}
	}
}

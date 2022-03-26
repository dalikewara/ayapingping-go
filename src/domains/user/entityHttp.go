package user

type HttpResponse struct {
	Error   bool        `json:"error"`
	Message string      `json:"message"`
	Data    interface{} `json:"data"`
}

type HttpCreateRequest struct {
	Username string `json:"username" binding:"required"`
}

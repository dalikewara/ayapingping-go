package user_example

// APIServiceLoginRequest object.
type APIServiceLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// APIServiceLoginResponse object.
type APIServiceLoginResponse struct {
	Code     string `json:"code" binding:"required"`
	Message  string `json:"message" binding:"required"`
	Username string `json:"username" binding:"required"`
}

// APIServiceLoginContextRequest object.
type APIServiceLoginContextRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

// APIServiceLoginContextResponse object.
type APIServiceLoginContextResponse struct {
	Code     string `json:"code" binding:"required"`
	Message  string `json:"message" binding:"required"`
	Username string `json:"username" binding:"required"`
}

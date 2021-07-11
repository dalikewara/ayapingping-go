package user_example

type APIServiceLoginRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type APIServiceLoginResponse struct {
	Code     string `json:"code" binding:"required"`
	Message  string `json:"message" binding:"required"`
	Username string `json:"username" binding:"required"`
}

type APIServiceLoginContextRequest struct {
	Username string `json:"username" binding:"required"`
	Password string `json:"password" binding:"required"`
}

type APIServiceLoginContextResponse struct {
	Code     string `json:"code" binding:"required"`
	Message  string `json:"message" binding:"required"`
	Username string `json:"username" binding:"required"`
}

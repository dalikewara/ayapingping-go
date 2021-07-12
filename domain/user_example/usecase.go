package user_example

import "context"

// UseCase interface object.
type UseCase interface {
	// Login logins user_example.
	Login(request UseCaseLoginRequest) UseCaseLoginResponse

	// LoginContext logins user_example with context.
	LoginContext(request UseCaseLoginContextRequest) UseCaseLoginContextResponse
}

// UseCaseLoginRequest object.
type UseCaseLoginRequest struct {
	Username string
	Password string
}

// UseCaseLoginResponse object.
type UseCaseLoginResponse struct {
	User  *Entity
	Error error
}

// UseCaseLoginContextRequest object.
type UseCaseLoginContextRequest struct {
	Ctx      context.Context
	Username string
	Password string
}

// UseCaseLoginContextResponse object.
type UseCaseLoginContextResponse struct {
	User  *Entity
	Error error
}

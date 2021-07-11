package user_example

import "context"

type UseCase interface {
	// Login logins user_example.
	Login(request UseCaseLoginRequest) UseCaseLoginResponse

	// LoginContext logins user_example with context.
	LoginContext(request UseCaseLoginContextRequest) UseCaseLoginContextResponse
}

type UseCaseLoginRequest struct {
	Username string
	Password string
}

type UseCaseLoginResponse struct {
	User  *Entity
	Error error
}

type UseCaseLoginContextRequest struct {
	Ctx      context.Context
	Username string
	Password string
}

type UseCaseLoginContextResponse struct {
	User  *Entity
	Error error
}

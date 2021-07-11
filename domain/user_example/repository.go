package user_example

import "context"

type Repository interface {
	// GetByUsernamePassword gets user_example data by username and password.
	GetByUsernamePassword(request RepositoryGetByUsernamePasswordRequest) RepositoryGetByUsernamePasswordResponse

	// GetByUsernamePasswordContext gets user_example data with context by username and password.
	GetByUsernamePasswordContext(request RepositoryGetByUsernamePasswordContextRequest) RepositoryGetByUsernamePasswordContextResponse
}

type RepositoryGetByUsernamePasswordRequest struct {
	Username string
	Password string
}

type RepositoryGetByUsernamePasswordResponse struct {
	User  *Entity
	Error error
}

type RepositoryGetByUsernamePasswordContextRequest struct {
	Ctx      context.Context
	Username string
	Password string
}

type RepositoryGetByUsernamePasswordContextResponse struct {
	User  *Entity
	Error error
}

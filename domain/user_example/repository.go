package user_example

import "context"

// Repository interface object.
type Repository interface {
	// GetByUsernamePassword gets user_example data by username and password.
	GetByUsernamePassword(request RepositoryGetByUsernamePasswordRequest) RepositoryGetByUsernamePasswordResponse

	// GetByUsernamePasswordContext gets user_example data with context by username and password.
	GetByUsernamePasswordContext(request RepositoryGetByUsernamePasswordContextRequest) RepositoryGetByUsernamePasswordContextResponse
}

// RepositoryGetByUsernamePasswordRequest object.
type RepositoryGetByUsernamePasswordRequest struct {
	Username string
	Password string
}

// RepositoryGetByUsernamePasswordResponse object.
type RepositoryGetByUsernamePasswordResponse struct {
	User  *Entity
	Error error
}

// RepositoryGetByUsernamePasswordContextRequest object.
type RepositoryGetByUsernamePasswordContextRequest struct {
	Ctx      context.Context
	Username string
	Password string
}

// RepositoryGetByUsernamePasswordContextResponse object.
type RepositoryGetByUsernamePasswordContextResponse struct {
	User  *Entity
	Error error
}

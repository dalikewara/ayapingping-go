package user

import "context"

type NewUseCaseParam struct {
	Service ServiceInterface
	// If you need to use other services from other domains, you can pass them here.
	// Example:
	// EmailService EmailServiceInterface
	// SMSService   SMSServiceInterface
}

type UseCaseGetAllParam struct {
	Ctx context.Context // If you want to use context.
}

type UseCaseGetAllResult struct {
	Users *[]User `json:"users"`
	Error error   `json:"error"`
}

type UseCaseCreateParam struct {
	Username string          `json:"username"`
	Ctx      context.Context // If you want to use context.
}

type UseCaseCreateResult struct {
	Error error `json:"error"`
}

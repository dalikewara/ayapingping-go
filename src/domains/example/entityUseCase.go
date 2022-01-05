package example

import "context"

// NewUseCaseParam is the argument sets of the NewUseCase function.
type NewUseCaseParam struct {
	ExampleService ServiceInterface
	// If you need to use other services, you can pass them here.
	// Example:
	// EmailService EmailServiceInterface
	// SMSService   SMSServiceInterface
}

// UseCaseGetAndChangeNameParam is the argument sets of the UseCaseInterface.GetAndChangeName method.
type UseCaseGetAndChangeNameParam struct {
	Id   int64
	Name string
	Ctx  *context.Context // If you want to use context.
}

// UseCaseGetAndChangeNameResult is the response model of the UseCaseInterface.GetAndChangeName method.
type UseCaseGetAndChangeNameResult struct {
	Example   *Example
	IsChanged bool
	Error     error
}

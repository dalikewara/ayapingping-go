package example

// NewUseCaseParam param.
type NewUseCaseParam struct {
	ExampleService ServiceInterface
	// If the use case need to use other services, you can pass them here.
	// Example:
	// EmailService EmailServiceInterface
	// SMSService SMSServiceInterface
}

// UseCaseGetAndChangeNameParam param.
type UseCaseGetAndChangeNameParam struct {
	Id int64
	Name string
	// Ctx context.Context  // If you want to use context.
}

// UseCaseGetAndChangeNameResult result.
type UseCaseGetAndChangeNameResult struct {
	Example *Example
	IsChanged bool
	Error error
}
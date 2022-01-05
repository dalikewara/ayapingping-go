package example

import "context"

// NewServiceParam is the argument sets of the NewService function.
type NewServiceParam struct {
	Repository RepositoryInterface
}

// ServiceGetParam is the argument sets of the ServiceInterface.Get method.
type ServiceGetParam struct {
	Id  int64
	Ctx *context.Context // If you want to use context.
}

// ServiceGetResult is the response model of the ServiceInterface.Get method.
type ServiceGetResult struct {
	Example *Example
	Error   error
}

// ServiceUpdateNameParam is the argument sets of the ServiceInterface.UpdateName method.
type ServiceUpdateNameParam struct {
	Id   int64
	Name string
	Ctx  *context.Context // If you want to use context.
}

// ServiceUpdateNameResult is the response model of the ServiceInterface.UpdateName method.
type ServiceUpdateNameResult struct {
	Error error
}

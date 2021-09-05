package example

// NewServiceParam parameter.
type NewServiceParam struct {
	Repository RepositoryInterface
}

// ServiceGetParam parameter.
type ServiceGetParam struct {
	Id int64
	// Ctx context.Context  // If you want to use context.
}

// ServiceGetResult result.
type ServiceGetResult struct {
	Example *Example
	Error error
}

// ServiceUpdateNameParam parameter.
type ServiceUpdateNameParam struct {
	Id int64
	Name string
	// Ctx context.Context  // If you want to use context.
}

// ServiceUpdateNameResult result.
type ServiceUpdateNameResult struct {
	Error error
}

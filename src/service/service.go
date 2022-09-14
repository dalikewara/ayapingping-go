package service

type Example interface {
	// GetAll gets all example data.
	GetAll(param ExampleGetAllParam) ExampleGetAllResult
}

type Service struct {
	Example Example
}

// New generates new service.
func New(param NewParam) *Service {
	panic("implement me")
}

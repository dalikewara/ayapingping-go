package repository

type Example interface {
	// FindAll finds all example data.
	FindAll(param ExampleFindAllParam) ExampleFindAllResult
}

type Repository struct {
	Example Example
}

// New generates new repository.
func New(param NewParam) *Repository {
	panic("implement me")
}

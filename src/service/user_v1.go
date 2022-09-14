package service

import (
	"github.com/dalikewara/ayapingping-go/v2/src/repository"
)

type example struct {
	exampleRepo repository.Example
}

// NewExample generates new example that implements Example.
func NewExample(param NewExampleParam) Example {
	panic("implement me")
}

// GetAll gets all example data.
func (s *example) GetAll(param ExampleGetAllParam) ExampleGetAllResult {
	panic("implement me")
}

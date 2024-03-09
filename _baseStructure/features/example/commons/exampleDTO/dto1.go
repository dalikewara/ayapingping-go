package exampleDTO

import "github.com/dalikewara/ayapingping-go/v4/_baseStructure/domain"

func NewDTO1(example *domain.Example) *domain.ExampleDTO1 {
	return &domain.ExampleDTO1{
		ID:        example.ID,
		Username:  example.Username,
		CreatedAt: example.CreatedAt,
	}
}

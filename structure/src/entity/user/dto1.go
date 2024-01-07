package user

import "github.com/dalikewara/ayapingping-go/v3/structure/src/entity/userProfile"

type DTO1 struct {
	Username string            `json:"username"`
	Profile  *userProfile.DTO1 `json:"profile"`
}

// NewDTO1 creates new data transfer object (form 1) for User entity
func NewDTO1(model *Model, profile *userProfile.DTO1) *DTO1 {
	if model == nil {
		return nil
	}

	return &DTO1{
		Username: model.Username,
		Profile:  profile,
	}
}

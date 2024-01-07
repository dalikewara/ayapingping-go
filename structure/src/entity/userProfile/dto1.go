package userProfile

type DTO1 struct {
	Name string `json:"name"`
}

// NewDTO1 creates new data transfer object (form 1) for User Profile entity
func NewDTO1(model *Model) *DTO1 {
	if model == nil {
		return nil
	}

	return &DTO1{
		Name: model.Name,
	}
}

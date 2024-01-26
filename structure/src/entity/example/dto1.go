package example

type DTO1 struct {
	ID   uint64 `json:"id"`
	Name string `json:"name"`
}

// NewDTO1 generates new data transfer object (form 1) from Model
func NewDTO1(model *Model) *DTO1 {
	if model == nil {
		return nil
	}

	return &DTO1{
		ID:   model.ID,
		Name: model.Name,
	}
}

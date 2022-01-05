package example

// Example is the main entity model of the domain.
type Example struct {
	Id   int64  `json:"id" bson:"_id,omitempty"`
	Name string `json:"name"`
}

// ValidateName validates Example.Name.
func (e *Example) ValidateName() error {
	return nil
}

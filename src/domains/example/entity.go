package example

// Example entity.
type Example struct {
	Id int64  `json:"id" bson:"_id,omitempty"`
	Name string `json:"name"`
}

// ValidateName validates `example` name.
func (e *Example) ValidateName() error {
	return nil
}

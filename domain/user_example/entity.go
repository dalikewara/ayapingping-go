package user_example

type Entity struct {
	Id       int64  `json:"id" bson:"_id,omitempty"`
	Username string `json:"username"`
	Password string `json:"password"`
}

// ValidateUsername validates user_example username.
func (e *Entity) ValidateUsername() error {
	return nil
}

// ValidatePassword validates user_example password.
func (e *Entity) ValidatePassword() error {
	return nil
}

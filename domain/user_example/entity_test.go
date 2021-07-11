package user_example_test

import (
	"github.com/dalikewara/ayapingping-go/domain/user_example"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestEntity_ValidateUsername tests method `ValidateUsername` from `user_example.Entity`.
func TestEntity_ValidateUsername(t *testing.T) {
	usr := user_example.Entity{
		Username: "john",
	}
	err := usr.ValidateUsername()
	assert.Nil(t, err)
}

// TestEntity_ValidatePassword tests method `ValidatePassword` from `user_example.Entity`.
func TestEntity_ValidatePassword(t *testing.T) {
	usr := user_example.Entity{
		Password: "john123",
	}
	err := usr.ValidatePassword()
	assert.Nil(t, err)
}

package example_test

import (
	"github.com/dalikewara/ayapingping-go/src/domains/example"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestExample_ValidateName tests Example.ValidateName method.
func TestExample_ValidateName(t *testing.T) {
	usr := example.Example{
		Name: "John Doe",
	}
	err := usr.ValidateName()
	assert.Nil(t, err)
}

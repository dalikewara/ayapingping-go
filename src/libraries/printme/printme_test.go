package printme_test

import (
	"github.com/dalikewara/ayapingping-go/src/libraries/printme"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestPrint tests printme.Print method.
func TestPrint(t *testing.T) {
	me := printme.Print("dali kewara")
	assert.Equal(t, "dali kewara", me)
}

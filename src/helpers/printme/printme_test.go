package printme_test

import (
	"github.com/dalikewara/ayapingping-go/src/helpers/printme"
	"testing"
)

// TestPrint tests `printme.Print` method.
func TestPrint(t *testing.T) {
	printme.Print("john doe")
}

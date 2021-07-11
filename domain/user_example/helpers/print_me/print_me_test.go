package print_me_test

import (
	"github.com/dalikewara/ayapingping-go/domain/user_example/helpers/print_me"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestPrintMe(t *testing.T) {
	text := print_me.PrintMe("john doe")
	assert.Equal(t, "john doe", text)
}

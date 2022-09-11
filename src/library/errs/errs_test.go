package errs_test

import (
	"github.com/dalikewara/ayapingping-go/src/library/errs"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNew test New function.
func TestNew(t *testing.T) {
	err := errs.New("01", "data not found")
	assert.NotNil(t, err)
	assert.Implements(t, (*errs.Errs)(nil), err)
}

// TestNewWithHttpStatus test NewWithHttpStatus function.
func TestNewWithHttpStatus(t *testing.T) {
	err := errs.NewWithHttpStatus("01", "data not found", 200)
	assert.NotNil(t, err)
	assert.Implements(t, (*errs.Errs)(nil), err)
}

// TestErrs_GetCode test errs.GetCode method.
func TestErrs_GetCode(t *testing.T) {
	err := errs.New("01", "data not found")
	assert.NotNil(t, err)
	assert.Equal(t, "01", err.GetCode())
}

// TestErrs_GetMessage test errs.GetMessage method.
func TestErrs_GetMessage(t *testing.T) {
	err := errs.New("01", "data not found")
	assert.NotNil(t, err)
	assert.Equal(t, "data not found", err.GetMessage())
}

// TestErrs_GetHttpStatus test errs.GetHttpStatus method.
func TestErrs_GetHttpStatus(t *testing.T) {
	t.Run("http status 0", func(t *testing.T) {
		err := errs.New("01", "data not found")
		assert.NotNil(t, err)
		assert.Equal(t, 0, err.GetHttpStatus())
	})
	t.Run("OK", func(t *testing.T) {
		err := errs.NewWithHttpStatus("01", "data not found", 200)
		assert.NotNil(t, err)
		assert.Equal(t, 200, err.GetHttpStatus())
	})
}

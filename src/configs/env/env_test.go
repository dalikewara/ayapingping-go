package env_test

import (
	"github.com/dalikewara/ayapingping-go/src/configs/env"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestGenerate tests `env.Generate` method.
func TestGenerate(t *testing.T) {
	cfg, err := env.Generate()
	assert.Nil(t, err)
	assert.Equal(t, env.Variable{}, cfg)
}

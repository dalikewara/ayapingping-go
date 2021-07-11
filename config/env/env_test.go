package env_test

import (
	"github.com/dalikewara/ayapingping-go/config/env"
	"github.com/stretchr/testify/assert"
	"testing"
)

// TestNewEnv tests `env.NewEnv` method.
func TestNewEnv(t *testing.T) {
	cfg, err := env.NewEnv()
	assert.Nil(t, err)
	assert.Equal(t, env.Config{}, cfg)
}

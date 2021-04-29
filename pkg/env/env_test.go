package env_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/pkg/env"
	"github.com/stretchr/testify/assert"
)

func Test_Load(t *testing.T) {
	err := env.Load(".env.testing")
	assert.NoError(t, err)
}

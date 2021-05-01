package bcrypt_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/pkg/bcrypt"
	"github.com/stretchr/testify/assert"
)

func TestHashPasswordAndCheckPasswordHash(t *testing.T) {
	tc := []struct {
		name     string
		password string
		expected bool
	}{
		{"one", "password", true},
	}

	for _, tt := range tc {
		t.Run(tt.name, func(t *testing.T) {
			act, err := bcrypt.HashPassword(tt.password)
			assert.NoError(t, err)
			assert.NotEqual(t, "", act)
			exp := bcrypt.CheckPasswordHash(tt.password, act)
			assert.Equal(t, tt.expected, exp)
		})
	}
}

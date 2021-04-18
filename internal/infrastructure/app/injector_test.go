package app_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/app"
	"github.com/stretchr/testify/assert"
)

func Test_Injector(t *testing.T) {
	act := app.Injection(nil)
	assert.NotNil(t, act)
}

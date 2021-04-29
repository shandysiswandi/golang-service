package app_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/app"
	"github.com/stretchr/testify/assert"
)

func TestSuccess(t *testing.T) {
	act := app.Success("message", nil)
	ref := act.(app.SuccessBody)

	assert.Equal(t, false, ref.Error)
	assert.Equal(t, "message", ref.Message)
	assert.Equal(t, nil, ref.Data)
}

func TestSuccessForTest(t *testing.T) {
	body := `{"error":false,"message":"message","data":null}`
	act, err := app.SuccessForTest(body)

	assert.NoError(t, err)
	assert.Equal(t, false, act.Error)
	assert.Equal(t, "message", act.Message)
	assert.Equal(t, nil, act.Data)

	body = "error"
	act, err = app.SuccessForTest(body)
	assert.Error(t, err)
	assert.Empty(t, act)
}

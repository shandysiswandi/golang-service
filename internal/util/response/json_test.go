package response_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/internal/util/response"
	"github.com/stretchr/testify/assert"
)

func Test_Success(t *testing.T) {
	act := response.Success("message", nil)
	ref := act.(response.SuccessBody)

	assert.Equal(t, false, ref.Error)
	assert.Equal(t, "message", ref.Message)
	assert.Equal(t, nil, ref.Data)
}

func Test_SuccessForTest(t *testing.T) {
	body := `{"error":false,"message":"message","data":null}`
	act, err := response.SuccessForTest(body)

	assert.NoError(t, err)
	assert.Equal(t, false, act.Error)
	assert.Equal(t, "message", act.Message)
	assert.Equal(t, nil, act.Data)

	body = "error"
	act, err = response.SuccessForTest(body)
	assert.Error(t, err)
	assert.Empty(t, act)
}

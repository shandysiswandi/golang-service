package handler_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/rpc/handler"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	h := handler.New()
	assert.NotNil(t, h)
}

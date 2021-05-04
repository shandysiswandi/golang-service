package handler_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/app/handler"
	"github.com/stretchr/testify/assert"
)

func TestNew(t *testing.T) {
	h := handler.New(handler.HandlerConfig{})
	assert.NotNil(t, h)
}

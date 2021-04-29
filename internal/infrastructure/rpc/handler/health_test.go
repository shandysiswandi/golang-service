package handler_test

import (
	"context"
	"testing"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/rpc/handler"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/rpc/pb"
	"github.com/stretchr/testify/assert"
)

func TestHealth(t *testing.T) {
	ctx := context.Background()
	h := handler.New()
	act, err := h.Health(ctx, &pb.Empty{})

	assert.NoError(t, err)
	assert.NotNil(t, act)
	assert.Equal(t, "this service is healthy", act.GetMessage())
}

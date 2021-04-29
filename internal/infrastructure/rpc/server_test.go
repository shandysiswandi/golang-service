package rpc_test

import (
	"testing"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/rpc"
	"github.com/stretchr/testify/assert"
)

func TestNewServer(t *testing.T) {
	srv := rpc.NewServer()
	assert.NotNil(t, srv)
}

package rpc

import "github.com/shandysiswandi/echo-service/internal/infrastructure/rpc/handler"

func NewServer() *handler.Handler {
	return handler.New()
}

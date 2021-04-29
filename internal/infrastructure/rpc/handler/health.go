package handler

import (
	"context"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/rpc/pb"
)

func (h *Handler) Health(context.Context, *pb.Empty) (*pb.Message, error) {
	msg := &pb.Message{Message: "this service is healthy"}
	return msg, nil
}

package handler

import (
	"context"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/rpc/pb"
)

func (h *Handler) Home(context.Context, *pb.Empty) (*pb.Message, error) {
	msg := &pb.Message{Message: "welcome home"}
	return msg, nil
}

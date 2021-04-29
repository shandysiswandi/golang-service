package handler

import (
	"context"
	"time"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/rpc/pb"
)

func (h *Handler) Graceful(context.Context, *pb.Empty) (*pb.Message, error) {
	time.Sleep(1 * time.Second)
	msg := &pb.Message{Message: "graceful"}
	return msg, nil
}

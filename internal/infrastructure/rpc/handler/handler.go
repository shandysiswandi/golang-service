package handler

import "github.com/shandysiswandi/echo-service/internal/infrastructure/rpc/pb"

type Handler struct {
	pb.UnimplementedMessageServiceServer
}

func New() *Handler {
	return &Handler{}
}

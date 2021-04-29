package main

import (
	"log"
	"net"

	"github.com/shandysiswandi/echo-service/internal/infrastructure/rpc"
	"github.com/shandysiswandi/echo-service/internal/infrastructure/rpc/pb"
	"google.golang.org/grpc"
)

func main() {
	listen, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatal(err)
	}

	sHandler := rpc.NewServer()
	sRPC := grpc.NewServer()

	pb.RegisterMessageServiceServer(sRPC, sHandler)

	println("Server gRPC running on => 0.0.0.0:50051")
	if err := sRPC.Serve(listen); err != nil {
		log.Fatal(err)
	}
}

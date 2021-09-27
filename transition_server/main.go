package main

import (
	"context"
	"log"
	"net"
	pb "probuf_reflect_test/transition"

	"google.golang.org/grpc"
)

const port = ":50052"

type serverImpl struct {
	pb.UnimplementedTransitionServer
}

func (s *serverImpl) Simple(ctx context.Context, param *pb.Param) (*pb.Res, error) {
	return &pb.Res{Value: 52}, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error occurs: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTransitionServer(grpcServer, &serverImpl{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

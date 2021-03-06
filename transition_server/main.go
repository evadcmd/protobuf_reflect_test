package main

import (
	"context"
	"io/ioutil"
	"log"
	"net"
	pb "probuf_reflect_test/transition"

	"google.golang.org/grpc"
	"google.golang.org/protobuf/encoding/protojson"
)

const port = ":50052"
const pathJsonSimple = "./res/simple.json"

type transitionImpl struct {
	pb.UnimplementedTransitionServer
}

func (s *transitionImpl) Simple(ctx context.Context, param *pb.Param) (*pb.Res, error) {
	jsonb, err := ioutil.ReadFile(pathJsonSimple)
	if err != nil {
		log.Fatalf("error occurs when parsing json file: %v", err)
	}
	res := pb.Res{}
	if err = protojson.Unmarshal(jsonb, &res); err != nil {
		log.Fatalf("error occurs when unmarshaling json data: %v", err)
	}
	return &res, nil
}

func main() {
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("error occurs: %v", err)
	}
	grpcServer := grpc.NewServer()
	pb.RegisterTransitionServer(grpcServer, &transitionImpl{})
	log.Printf("server listening at %v", lis.Addr())
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

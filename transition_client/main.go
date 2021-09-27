package main

import (
	"context"
	"fmt"
	"log"

	pb "probuf_reflect_test/transition"

	"google.golang.org/grpc"
)

const port = ":50052"

func main() {
	conn, err := grpc.Dial(port, grpc.WithInsecure())
	if err != nil {
		log.Fatalf("error occurs: %v", err)
	}
	client := pb.NewTransitionClient(conn)
	res, err := client.Simple(context.TODO(), &pb.Param{})
	if err != nil {
		log.Fatalf("error occurs when calling a rpc method: %v", err)
	}
	fmt.Println(res)
}

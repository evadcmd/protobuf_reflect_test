package main

import (
	"context"
	"fmt"
	pb "probuf_reflect_test/transition"

	"google.golang.org/grpc"
)

func main() {
	serviceDescriptors := pb.File_transition_proto.Services()
	for i := 0; i < serviceDescriptors.Len(); i++ {
		serviceDescriptor := serviceDescriptors.Get(i)
		methodDescs := []grpc.MethodDesc{}
		methodDescriptors := serviceDescriptor.Methods()
		for j := 0; j < methodDescriptors.Len(); j++ {
			methodDescriptor := methodDescriptors.Get(j)
			methodDescs = append(methodDescs, grpc.MethodDesc{
				MethodName: string(methodDescriptor.FullName()),
				Handler: func(srv interface{}, ctx context.Context, dec func(interface{}) error, interceptor grpc.UnaryServerInterceptor) (interface{}, error) {
					return nil, nil
				},
			})
			fmt.Println(methodDescriptor.FullName())
		}
		serviceDesc := grpc.ServiceDesc{
			Methods: methodDescs,
		}
		fmt.Println(serviceDesc)
	}
}

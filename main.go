package main

import (
	"context"
	"fmt"
	"time"

	"github.com/fullstorydev/grpcurl"
	"github.com/jhump/protoreflect/grpcreflect"
	"google.golang.org/grpc"
	reflectpb "google.golang.org/grpc/reflection/grpc_reflection_v1alpha"
)

var (
	defaultReplicas int32 = 1
	sourceReflect   grpcurl.DescriptorSource
	ccReflect       *grpc.ClientConn
)

func main() {
	fmt.Println("Hello, world.")

	var portReflect int = 80
	var host = "10.110.84.49"
	var err error
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	println(fmt.Sprintf("** performing DialContext at %s:%d", host, portReflect))

	ccReflect, err = grpc.DialContext(ctx, fmt.Sprintf("%s:%d", host, portReflect), grpc.WithInsecure(), grpc.WithBlock(), grpc.WithTimeout(2*time.Second))
	if err == nil {
		println("** Got a respons")
		println(ccReflect)

		defer ccReflect.Close()
		refClient := grpcreflect.NewClient(context.Background(), reflectpb.NewServerReflectionClient(ccReflect))
		defer refClient.Reset()

		sourceReflect = grpcurl.DescriptorSourceFromServer(context.Background(), refClient)
		names, err := grpcurl.ListServices(sourceReflect)
		if err != nil {
			panic(err)
		}
		println(fmt.Sprintf("list of services: %s", names))
	} else {
		panic(err)
	}

}

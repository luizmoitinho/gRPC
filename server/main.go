package main

import (
	"context"
	"gRPC/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Hello(ctx context.Context, request *pb.HelloRequest) (*pb.HelloResponse, error) {
	result := "Hello" + request.GetName()

	//create response
	response := &pb.HelloResponse{
		Msg: result,
	}

	return response, nil
}

func main() {
	listener, err := net.Listen("tcp", "0.0.0.0:5001")

	if err != nil {
		log.Fatalf("Failed to listen %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterHelloServiceServer(grpcServer, &server{})

	//atribui e compara o valor contido em err
	if err := grpcServer.Serve(listener); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}

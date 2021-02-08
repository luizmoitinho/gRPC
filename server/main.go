package server

import (
	"gRPC/pb"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func main() {
	listener, err := net.Listen("tpc", "0.0.0.0:5001")

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

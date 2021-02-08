package main

import (
	"context"
	"gRPC/pb"
	"log"

	"google.golang.org/grpc"
)

func main() {

	connection, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Could not connect: %v", err)
	}
	//close connection pool
	defer connection.Close()

	client := pb.NewHelloServiceClient(connection)
	Hello(err, client)

}

func Hello(err error, client pb.HelloServiceClient) {
	request := &pb.HelloRequest{
		Name: "Luiz Moitinho",
	}

	//execute remote function
	res, err := client.Hello(context.Background(), request)

	if err != nil {
		log.Fatalf("Error during the executio: %vn", err)
	}
	log.Println(res.Msg)
}

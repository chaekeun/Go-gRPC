package main

import (
	"context"
	"fmt"
	"log"

	pb "github.com/chaekeun/Go-gRPC/lec-07-prg-04-serverstreaming/serverstreaming"
	"google.golang.org/grpc"
)

// go gRPC code use interf&struct(ex: ServerStreaminClient intf) instead of 'stub' in python
func recvMessage(client pb.ServerStreamingClient) {
	// GetServerResponse func get server response
	request := &pb.Number{Value: 5}
	stream, err := client.GetServerResponse(context.Background(), request)
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}
	waitc := make(chan struct{})

	// infinite loop to keep msg recv
	for {
		// blocking method waits till server send msg
		response, err := stream.Recv()
		if err != nil {
			// channel close
			close(waitc)
			return
		}
		fmt.Printf("[server to client] %s\n", response.Message)

	}
	stream.CloseSend()
	<-waitc
}

func run() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewServerStreamingClient(conn)
	recvMessage(client)
}

func main() {
	run()
}

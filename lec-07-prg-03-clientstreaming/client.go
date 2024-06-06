package main

import (
	"context"
	"fmt"
	"log"
	"time"

	pb "github.com/chaekeun/Go-gRPC/lec-07-prg-03-clientstreaming/clientstreaming"
	"google.golang.org/grpc"
)

func makeMessage(message string) *pb.Message {
	return &pb.Message{
		Message: message,
	}
}

func generateMessages() []*pb.Message {
	return []*pb.Message{
		makeMessage("message #1"),
		makeMessage("message #2"),
		makeMessage("message #3"),
		makeMessage("message #4"),
		makeMessage("message #5"),
	}
}

// go gRPC code use interf&struct instead of 'stub' in python
func sendMessage(client pb.ClientStreamingClient) {
	// GetServerResponse func get server response stream with ctx
	stream, err := client.GetServerResponse(context.Background())
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}

	messages := generateMessages()
	for _, msg := range messages {
		fmt.Printf("[client to server] %s\n", msg.Message)
		if err := stream.Send(msg); err != nil {
			log.Fatalf("could not send message: %v", err)
		}
		time.Sleep(time.Second)
	}
	response, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("could not receive response: %v", err)
	}
	fmt.Printf("[server to client] %d\n", response.Value)
}

func run() {
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewClientStreamingClient(conn)
	sendMessage(client)
}

func main() {
	run()
}

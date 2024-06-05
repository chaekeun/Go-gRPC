package main

import(
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/chaekeun/Go-gRPC/lec-07-prg-02-bidirectional-streaming/bidirectional"
)

func makeMessage(message string) *pb.Message {
	return &pb.Message{
		Message: message,
	}
}

func generateMessages() []*pb.Message {
	message := []*pb.Message{
		makeMessage("message #1"),
		makeMessage("message #2"),
		makeMessage("message #3"),
		makeMessage("message #4")
		makeMessage("message #5")
	}
	return messages
}

func sendMessage(client pb.BidrectionalClient) {
	stream, err := client.GetServerResponse(context.Background())
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}
	waitc := make(chan struct{})

	go func(){
		for {
			response, err := stream.Recv()
			if err != nil{
				close(waitc)
				return
			}
			fmt.Printf("[server to client] %s\n", response.Message)
		}
	}()

	for _, msg := range generateMessage(){
		fmt.Printf("[client to server] %s\n", msg.Message)
		if err := stream.Send(msg); err != nil {
			log.Fatalf("could not send message: %v", err)
		}
		time.Sleep(time.Second)
	}
	stream.CloseSend()
	<-waitc
}

func main(){
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	client := pb.NewBidirectionalClient(conn)
	sendMessage(client)
}
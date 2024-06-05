package main

import(
	"context"
	"fmt"
	"log"
	"time"

	"google.golang.org/grpc"
	pb "github.com/chaekeun/Go-gRPC/lec-07-prg-02-bidirectional-streaming/bidirectional"
)

// make message grpc obj using *pb.Message
func makeMessage(message string) *pb.Message {
	return &pb.Message{
		Message: message,
	}
}

func generateMessages() []*pb.Message {
	messages := []*pb.Message{
		makeMessage("message #1"),
		makeMessage("message #2"),
		makeMessage("message #3"),
		makeMessage("message #4"),
		makeMessage("message #5"),
	}
	return messages
}

// streaming communication using grpc client
func sendMessage(client pb.BidirectionalClient) {
	stream, err := client.GetServerResponse(context.Background())
	if err != nil {
		log.Fatalf("could not create stream: %v", err)
	}
	waitc := make(chan struct{})

	// using goroutine : async function for real time res
	go func(){
		// infinte loop to keep msg recv
		for { 
			// blocking method waits till server send msg
			response, err := stream.Recv()
			if err != nil{
				// channel close
				close(waitc)
				return
			}
			fmt.Printf("[server to client] %s\n", response.Message)
		}
	}()

	for _, msg := range generateMessages(){
		fmt.Printf("[client to server] %s\n", msg.Message)
		if err := stream.Send(msg); err != nil {
			log.Fatalf("could not send message: %v", err)
		}
		time.Sleep(time.Second)
	}
	stream.CloseSend()
	<-waitc
}

// send msg to server
func main(){
	// set grpc cahnnel
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()

	// make new bidriectional client
	client := pb.NewBidirectionalClient(conn)
	// call sendmessage function : send&recv using goroutine = concurrency
	sendMessage(client)
}
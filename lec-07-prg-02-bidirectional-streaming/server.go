package main

import (
	"fmt"
	"log"
	"net"
	"sync"

	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
	pb "github.com/chaekeun/Go-gRPC/lec-07-prg-02-bidirectional-streaming"
)

// go use struct to implent gRPC service.
type BidrectionalService struct{
	// embed default method
	pb.UnimplementedBidirectionalServer
}

// recv streaming req from client and res
func (s *BidirectionalService) GetServerResponse(stream pb.Bidirectional_GetServerResponseServer) error {
	fmt.Println("Server processing gRPC bidirectional streaming.")
	for {
		message, err := stream.Recv()
		if err != nil{
			return err
		}
		p, _ := peer.FromContext(stream.Context())
		fmt.Printf("Received message from %v: %s\n", p.Addr, meesage.Message)
		if err  := stream.Send(message); err != nil{
			return err
		}
	}
}

// set&start gRPC server 
func serve(){
	// tcp listner
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	
	//make new grpc server
	server := grpc.NewServer()

	// register bidrectionalservice struct into grpc server
	pb.RegisterBidirectionalServer(server, &BidirectionalService{})

	fmt.Println("Starting server. Listening on port 50051. ")
	if err := server.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}

func main(){
	serve()
}
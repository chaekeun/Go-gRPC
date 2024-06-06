package main

import (
	"fmt"
	"io"
	"log"
	"net"

	pb "github.com/chaekeun/Go-gRPC/lec-07-prg-03-clientstreaming/clientstreaming"
	"google.golang.org/grpc"
	"google.golang.org/grpc/peer"
)

type ClientStreamingServicer struct {
	pb.UnimplementedClientStreamingServer
}

func (s *ClientStreamingServicer) GetServerResponse(stream pb.ClientStreaming_GetServerResponseServer) error {
	fmt.Println("Server processing gRPC client-streaming.")
	count := 0
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			// When client stream ends, send the final response(=count) and close
			return stream.SendAndClose(&pb.Number{Value: int32(count)})
		}
		if err != nil {
			return err
		}
		count++
		p, _ := peer.FromContext(stream.Context())
		fmt.Printf("Received message from %v: %s\n", p.Addr, message.Message)
	}
}

func serve() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := grpc.NewServer()
	pb.RegisterClientStreamingServer(server, &ClientStreamingServicer{})

	fmt.Println("Starting server. Listening on port 50051.")
	if err := server.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

func main() {
	serve()
}

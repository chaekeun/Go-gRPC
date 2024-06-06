package main

import (
	"fmt"
	"log"
	"net"
	"time"

	pb "github.com/chaekeun/Go-gRPC/lec-07-prg-04-serverstreaming/serverstreaming"
	"google.golang.org/grpc"
)

type ServerStreamingServicer struct {
	pb.UnimplementedServerStreamingServer
}

func makeMessage(message string) *pb.Message {
	return &pb.Message{
		Message: message,
	}
}

// *pb.Request가 아니라 *pb.Number 아닌가?
func (s *ServerStreamingServicer) GetServerResponse(req *pb.Number, stream pb.ServerStreaming_GetServerResponseServer) error {
	messages := []*pb.Message{
		makeMessage("message #1"),
		makeMessage("message #2"),
		makeMessage("message #3"),
		makeMessage("message #4"),
		makeMessage("message #5"),
	}
	fmt.Printf("Server processing gRPC server-streaming {%d}.\n", req.Value)
	for _, message := range messages {
		if err := stream.Send(message); err != nil {
			return err
		}
		time.Sleep(1 * time.Second) // Simulate some processing time
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterServerStreamingServer(s, &ServerStreamingServicer{})
	fmt.Println("Starting server. Listening on port 50051.")
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

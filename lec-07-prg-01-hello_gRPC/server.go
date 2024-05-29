package main

// (2) import protoc struct
import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "github.com/chaekeun/Go-gRPC/lec-07-prg-01-hello_gRPC/helloGrpc"

// (3) original remotely called functions
//	hello_grpc "github.com/chaekeun/Go-gRPC/lec-07-prg-01-hello_gRPC/helloGrpc/hello_grpc"
)


// (4) using servicer struct created by protoc
type myServiceServer struct {
	pb.UnimplementedMyServiceServer
}

// (5) remote call rpc func
	// (5.1) user defined rpc function MyFunction
func (s *myServiceServer) MyFunction(ctx context.Context, req *pb.Mynumber) (*pb.MyNumber, error){
	// (5.2) user defined msg class
	res := &pb.MyNumber{
		// (5.3) pass input param to user defined rpc function and save return value
		Value: myFunc(req.GetValue()),
	}
	return res, nil
}

func main(){
// (8) open port 50051
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
// (6) make grpc server
	s := grpc.NewServer()

// (7) add (4)server by using func created by protoc
	pb.RegisterMyServiceServer(s, &myServiceServer{})
	
	log.Println("Starting server. Listening on port 50051.")
	if err := s.Serve(lis); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}

//(9) try execpt to maintain grpc.NewServer?

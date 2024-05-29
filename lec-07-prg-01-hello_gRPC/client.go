import (
        "context"
        "net"

// (1) import grpc module
        "google.golang.org/grpc"

// (2) import protoc struct
	helloGrpc "github.com/chaekeun/Go-gRPC/lec-07-prg-01-hello_gRPC/helloGrpc"

)

func main(){
// (3) grpc channel
	conn, err := grpc.Dial("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil{
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	
// (4) using (3)channel, create clientStub
	client := helloGrpc.NewMyServiceClient(conn)

// (5)ãmake msg to send remote func using protoc defined msg type
	request := &helloGrpc.MyNumber{Value: 4}

// (6) remote function call using clientStub
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	
	response, err := client.MyFunction(ctx, request)
	if err != nil{
		log.Fatalf("could not call MyFunction: %v", err)
	}
	log.Printf("gRPC result: %d", response.GetValue())
}

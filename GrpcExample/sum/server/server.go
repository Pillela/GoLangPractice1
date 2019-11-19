package main

import (
	"GrpcExample/sum/sumpb"
	"context"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct{}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest) (*sumpb.SumResponse, error) {
	fmt.Printf("Sum function was invoked with %v", req)
	value1 := req.GetA()
	value2 := req.GetB()
	result := value1 + value2
	res := &sumpb.SumResponse{
		Sumresult: result,
	}
	return res, nil
}

func main() {
	fmt.Println("Hello this is server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterAdditionServer(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Failed to server: %v", err)
	}
}

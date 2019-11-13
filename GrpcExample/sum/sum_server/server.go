package main

import (
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
}

func (*server) Sum(ctx context.Context, req *sumpb.SumRequest)

func main() {
	fmt.Println("This is server!")

	lis, err := net.Listen("tcp", "0.0.0.0:50051")
	if err != nil {
		log.Fatalf("Server Failed to listen: %v", err)
	}

	s := grpc.NewServer()
	sumpb.RegisterAdditionSever(s, &server{})

	if err := s.Serve(lis); err != nil {
		log.Fatalf("Server failed to serve: %v", err)
	}

}

package main

import (
	"GrpcExample/sum/sumpb"
	"context"
	"fmt"
	"log"

	"google.golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Issue with client connection: %v", err)
	}
	defer cc.Close()

	c := sumpb.NewAdditionClient(cc)

	var local_a int32
	fmt.Println("Enter the value of a")
	fmt.Scanln(&local_a)
	var local_b int32
	fmt.Println("Enter the value of b")
	fmt.Scanln(&local_b)

	req := &sumpb.SumRequest{
		A: local_a,
		B: local_b,
	}

	res, err := c.Sum(context.Background(), req)
	if err != nil {
		log.Fatalf("Error while calling Sum RPC:%v", err)
	}
	log.Printf("Response from Sum: %v", res.Sumresult)

}

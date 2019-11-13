package main

import (
	"GrpcExample/sum/sumpb"
	"fmt"
	"log"

	grpc "google/golang.org/grpc"
)

func main() {
	fmt.Println("Hello I'm Client")
	cc, err := grpc.Dial("localhost:50051", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("Issue with client connection: %v", err)
	}
	defer cc.close()

	c := sumpb.NewAdditionClient(cc)

	var local_a int
	fmt.Println("Enter the value of a")
	fmt.Scanln(&local_a)
	var local_b int
	fmt.Println("Enter the value of b")
	fmt.Scanln(&local_b)

	req:= greetpb.SumRequest{
		a: local_a;
		b: local_b;
	}

	

}

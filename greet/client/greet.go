package main

import (
	"context"
	"log"

	pb "github.com/shubhankars135/Go-GRPC-Unary-sample/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("Dogreet was invoked")

	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Shubhanker",
	})


	if err != nil {
		log.Fatalf("Could'nu %v\n", err)
	}

	log.Printf("Greeting %s\n", res.Result)
}
package main

import (
	"context"
	"log"

	pb "github.com/shubhankars135/Go-GRPC-Unary-sample/calculator/proto"
)

func doSum(c pb.CalculatorClient) {
	log.Println("Calculator was was invoked")

	res, err := c.Sum(context.Background(), &pb.SumRequest{
		FirstNum: 12,
		SecondNum: 15,
	})


	if err != nil {
		log.Fatalf("Could'nu %v\n", err)
	}

	log.Printf("Result is %d\n", res.Result)
}
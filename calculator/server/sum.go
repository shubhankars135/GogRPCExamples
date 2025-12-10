
package main

import (
	"context"
	"log"

	pb  "github.com/shubhankars135/Go-GRPC-Unary-sample/calculator/proto"
)

func (s *Server) Sum(ctx context.Context, in *pb.SumRequest) (*pb.SumResponse, error){
	log.Printf("Greet was invoked with %v\n", in)
	result := in.FirstNum + in.SecondNum
	return &pb.SumResponse{Result: result}, nil
}

package main

import (
	"fmt"
	"log"

	pb "github.com/shubhankars135/Go-GRPC-Unary-sample/greet/proto"
)

func (s *Server) GreetManyTimes(in *pb.GreetRequest, stream pb.GreetService_GreetManyTimesServer) error {
	log.Printf("Manytimes was invoked")

	for i := 0 ; i < 10; i ++ {
		res := fmt.Sprintf("Incoming request %s  %d " ,in.FirstName, i)
	

	stream.Send(&pb.GreetResponse{
		Result: res,
	})
	}
	return nil
}

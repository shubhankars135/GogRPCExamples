package main

import (
	"io"
	"log"

	pb "github.com/shubhankars135/Go-GRPC-Unary-sample/greet/proto"
)

func (s *Server) GreetEveryOne(stream pb.GreetService_GreetEveryOneServer) error {
	log.Println("GreetEveryone was invoked")

	for {
		req, err := stream.Recv()

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"

		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})
		if err != nil {
			log.Fatalf("Error %v\n", err)
		}
	}

}

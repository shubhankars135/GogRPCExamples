package main

import (
	"log"
	"net"
	"google.golang.org/grpc"

	pb "github.com/shubhankars135/Go-GRPC-Unary-sample/greet/proto"
)

var address string = "0.0.0.0:8001"

type Server struct {
	pb.GreetServiceServer
}

func main () {
	lis, err := net.Listen("tcp", address)

	if err != nil {
		log.Fatalf("Failed to listen on %v\n", err)
	}
	log.Printf("Listening on %s\n", address)

	s := grpc.NewServer()
	pb.RegisterGreetServiceServer(s, &Server{})



	if err = s.Serve(lis); err != nil {
		log.Fatalf("Failed to serve on %v\n", err)
	}

}
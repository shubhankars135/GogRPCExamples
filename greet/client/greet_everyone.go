package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/shubhankars135/Go-GRPC-Unary-sample/greet/proto"
)


func doGreetEveryOne(c pb.GreetServiceClient){
	log.Printf("Greet evervryone was invoked")

	stream, err := c.GreetEveryOne(context.Background())

	if err != nil {
		log.Fatalf("Error %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Shubh"},
		{FirstName: "Nik"},
		{FirstName: "Test"},
	}

	waitc := make(chan struct{})

	go func() {
		for _, req := range reqs {
			log.Printf("Send request %v\n", req)
			stream.Send(req)
			time.Sleep(1 * time.Second)
		}
		stream.CloseSend()
	}()

	go func ()  {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Fatalf("Error %v\n", err)
				break
			}

			log.Printf("Received %v\n", res.Result)
		}
		// Close the Waitgroup
		close(waitc)
	}()

	// Will wait for wait grp to be closed
	<- waitc

}

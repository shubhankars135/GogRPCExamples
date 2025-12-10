package main
import (
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
	"log"

	pb "github.com/shubhankars135/Go-GRPC-Unary-sample/calculator/proto"
)

var address string = "0.0.0.0:8001"

func main() {
	conn, err := grpc.NewClient(
		address, 
		grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err!=nil{
		log.Fatalf("failed %v\n", err)
	}

	defer conn.Close()

	c := pb.NewCalculatorClient(conn)
	doSum(c)
}

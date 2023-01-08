package main

import (
	"context"
	"log"
	"time"

	pb "github.com/behroozz/grpc-microservice/greet/proto"
)

func doLongGreet(c pb.GreetServiceClient) {
	log.Printf("doLongGreet was invoked")

	reqs := []*pb.GreetRequest{
		{FirstName: "Behrooz"},
		{FirstName: "Hesam"},
		{FirstName: "Ahmad"},
	}

	stream, err := c.LongGreet(context.Background())

	if err != nil {
		log.Fatalf("Error calling LongGreet %v\n", err)
	}

	for _, req := range reqs {
		log.Printf("Sendig req %v\n", req)
		stream.Send(req)
		time.Sleep(1 * time.Second)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("Error while recieving response from LongGreet %v\n", err)
	}

	log.Printf("LongGreet: %s\n", res.Result)
}

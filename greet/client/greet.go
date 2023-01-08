package main

import (
	"context"
	"log"

	pb "github.com/behroozz/grpc-microservice/greet/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("do Greet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "Behrooz",
	})

	if err != nil {
		log.Fatalf("Could not greet %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}

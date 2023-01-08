package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/behroozz/grpc-microservice/greet/proto"
)

func doGreetEveryone(c pb.GreetServiceClient) {
	log.Println("doGreetEveryone was invoked")

	stream, err := c.GreetEveryone(context.Background())

	if err != nil {
		log.Fatalf("Error while calling GreetEveryone %v\n", err)
	}

	reqs := []*pb.GreetRequest{
		{FirstName: "Behrooz"},
		{FirstName: "Hesam"},
		{FirstName: "Ahmad"},
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

	go func() {
		for {
			res, err := stream.Recv()

			if err == io.EOF {
				break
			}

			if err != nil {
				log.Printf("Error while reading the stream %v\n", err)
				break
			}

			log.Printf("Recieved: %v\n", res.Result)
		}

		close(waitc)
	}()

	<-waitc
}

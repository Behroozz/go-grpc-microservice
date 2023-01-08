package main

import (
	"io"
	"log"

	pb "github.com/behroozz/grpc-microservice/greet/proto"
)

func (s *Server) GreetEveryone(stream pb.GreetService_GreetEveryoneServer) error {
	log.Printf("GreetEveryone was invoked with %v\n: ", stream)

	for {
		req, err := stream.Recv()

		log.Printf("Recieving %v\n", req)

		if err == io.EOF {
			return nil
		}

		if err != nil {
			log.Fatalf("Error while reading client stream %v\n", err)
		}

		res := "Hello " + req.FirstName + "!"
		err = stream.Send(&pb.GreetResponse{
			Result: res,
		})

		if err != nil {
			log.Fatalf("Error while sending data to client %v\n", err)
		}
	}
}

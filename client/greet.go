package main

import (
	"context"
	"log"

	pb "github.com/manoj-negi/grpc-microservices/proto"
)

func doGreet(c pb.GreetServiceClient) {
	log.Println("doGreet was invoked")
	res, err := c.Greet(context.Background(), &pb.GreetRequest{
		FirstName: "sonam",
	})

	if err != nil {
		log.Fatalf("coluld not greet: %v\n", err)
	}

	log.Printf("Greeting: %s\n", res.Result)
}

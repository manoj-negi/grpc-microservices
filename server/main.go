package main

import (
	"log"
	"net"

	pb "github.com/manoj-negi/grpc-microservices/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:5051"

type Server struct {
	pb.GreetServiceServer
}

func main() {

	list, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed : %v\n", err)
	}
	log.Printf("Listening server %s", addr)

	s := grpc.NewServer()

	pb.RegisterGreetServiceServer(s, &Server{})

	defer s.Stop()
	if err := s.Serve(list); err != nil {
		log.Fatalf("fail to serve: %v\n", err)
	}

}

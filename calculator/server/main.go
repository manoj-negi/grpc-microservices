package main

import (
	"log"
	"net"

	pb "github.com/manoj-negi/grpc-microservices/calculator/proto"
	"google.golang.org/grpc"
)

var addr string = "0.0.0.0:5051"

type Server struct {
	pb.CalculatorServiceServer
}

func main() {

	lis, err := net.Listen("tcp", addr)
	if err != nil {
		log.Fatalf("fail to listen %s", err)
	}

	s := grpc.NewServer()

	pb.RegisterCalculatorServiceServer(s, &Server{})
	defer s.Stop()

	if err := s.Serve(lis); err != nil {
		log.Fatalf("fail to server, %v\n", err)
	}

}

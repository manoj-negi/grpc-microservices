package main

import (
	"log"
	"net"
)

var addr string = "0:0:0:0:5051"

func main() {

	list, err := net.Listen("tcp", addr)

	if err != nil {
		log.Fatalf("failed : %v\n", err)
	}
}

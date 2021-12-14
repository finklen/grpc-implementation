package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/finklen/grpc-implementation/go/helloworld/helloworld"
)

var (
	port = flag.Int("port", 50051, "Server port")
)

func main() {
	flag.Parse()

	_, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	// _ := grpc.NewServer()

	pb.RegisterGreeterServer()
}

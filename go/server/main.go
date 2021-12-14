package main

import (
	"flag"
	"fmt"
	"log"
	"net"

	pb "github.com/finklen/grpc-implementation/go/helloworld/helloworld"
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterSever
}

func (s *server) SayHello(ctx content.Context, in ) {

}

var (
	port = flag.Int("port", 50051, "Server port")
)

func main() {
	flag.Parse()

	_, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", *port))
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	s := grpc.NewServer()

	pb.RegisterGreeterServer(s, &server{})
}

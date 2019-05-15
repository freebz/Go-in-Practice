// Listing 10.12  gRPC server responding to requests

package main

import (
	"log"
	"net"

	pb "github.com/Masterminds/go-in-practice/chapter10/hellopb"
	"golang.org/x/net/context"
	"google.golang.org/grpc"
)

type server struct{}

func (s *server) Say(ctx context.Context, \
 in *pb.HelloRequest) (*pb.HelloResponse, error) {
	msg := "Hello " + in.Name + "!"
	return &pb.HelloResponse(Message: msg), nil
}

func main() {
	l, err := net.Listen("tcp", ":55555")
	if err != nil {
		...
	}
	s := grpc.NewServer()
	pb.RegisterHelloServer(s, &server())
	s.Serve(l)
}

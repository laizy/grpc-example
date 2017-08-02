package main

import (
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-example/ex1-hello/proto"
	"log"
	"net"
)

type helloService struct{}

func (self *helloService) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

const Port = ":8000"

func main() {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterHelloServer(s, &helloService{})

	log.Printf("serve on \n", Port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

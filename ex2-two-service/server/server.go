package main

import (
	context "golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-example/ex2-two-service/proto"
	"log"
	"net"
)

type ex2Service struct{}

func (self *ex2Service) SayHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

func (self *ex2Service) Fib(ctx context.Context, req *pb.FibMsg) (*pb.FibMsg, error) {
	return &pb.FibMsg{Num: fib(req.Num)}, nil
}

func fib(num int64) int64 {
	if num == 0 {
		return 0
	} else if num == 1 {
		return 1
	} else {
		return fib(num-1) + fib(num-2)
	}

}

const Port = ":8000"

func main() {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterHelloServer(s, &ex2Service{})
	pb.RegisterFibonacciServer(s, &ex2Service{})

	log.Printf("serve on \n", Port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

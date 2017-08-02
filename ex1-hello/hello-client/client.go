package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-example/ex1-hello/proto"
	"log"
)

const Port = ":8000"

func main() {
	conn, err := grpc.Dial("localhost"+Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewHelloClient(conn)

	names := []string{"name1", "name2", "name3"}

	for _, name := range names {
		r, err := client.SayHello(context.Background(), &pb.HelloRequest{Name: name})
		if err != nil {
			log.Fatalf("call hello err: %v", err)
		}

		log.Println(r.Message)
	}

}

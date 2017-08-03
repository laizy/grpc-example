package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-example/ex3-client-stream/proto"
	"log"
)

const Port = ":8000"

func main() {
	conn, err := grpc.Dial("localhost"+Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	client := pb.NewUploaderClient(conn)

	stream, err := client.Upload(context.Background())
	if err != nil {
		log.Fatalf("failed to call Upload, %v", err)
	}

	filepart := []byte("hello")

	for i := 0; i < 40; i++ {
		err := stream.Send(&pb.UploadRequest{Offset: uint64(i * len(filepart)), Data: filepart})
		if err != nil {
			log.Fatalf("failed to upload data, %v", err)
		}
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("failed to recv data, %v", err)
	}
	log.Println(res.Message)

}

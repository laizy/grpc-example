package main

import (
	"bytes"
	"errors"
	"fmt"
	// context "golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-example/ex3-client-stream/proto"
	"io"
	"log"
	"net"
)

type uploaderService struct {
}

func (self *uploaderService) Upload(stream pb.Uploader_UploadServer) error {
	var data bytes.Buffer
	var offset uint64

	for {
		req, err := stream.Recv()
		if err == io.EOF {
			break
		} else if err != nil {
			log.Println("recv err", err)
			return err
		}
		if req.Offset != offset {
			msg := fmt.Sprintf("wrong request offset, need %v, recv %v", offset, req.Offset)
			log.Println(msg)
			return errors.New(msg)
		}
		offset += uint64(len(req.Data))
		data.Write(req.Data)
	}

	file := data.Bytes()

	log.Printf("recv file, size %d, content: %s...", offset, file[:60])

	err := stream.SendAndClose(&pb.UploadResponse{Message: "success"})
	if err != nil {
		log.Printf("recv file ")
	}

	return nil
}

const Port = ":8000"

func main() {
	lis, err := net.Listen("tcp", Port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()

	pb.RegisterUploaderServer(s, &uploaderService{})

	log.Printf("serve on %v\n", Port)

	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

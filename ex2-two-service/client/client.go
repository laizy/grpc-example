package main

import (
	"golang.org/x/net/context"
	"google.golang.org/grpc"
	pb "grpc-example/ex2-two-service/proto"
	"log"
	"strconv"
	"sync"
)

const Port = ":8000"

func main() {
	conn, err := grpc.Dial("localhost"+Port, grpc.WithInsecure())
	if err != nil {
		log.Fatalln(err)
	}
	defer conn.Close()

	//	实际测试发现同一个conn上创建两个Client并发发送请求，并没有出现并发数据读写问题，
	//  从生成的pb.go文件可以看出 fibclient 和helloclient 只是对conn的一个封装，
	//  对应的并发请求 均通过conn.Invoke 函数执行，所以不同的客户端并发调用和
	//  同一个客户端并发调用并没有区别
	helloclient := pb.NewHelloClient(conn)
	fibclient := pb.NewFibonacciClient(conn)

	var wait sync.WaitGroup

	wait.Add(1)
	go func() {
		defer wait.Done()

		for i := 0; i < 40000; i++ {
			name := "name" + strconv.Itoa(i)
			r, err := helloclient.SayHello(context.Background(), &pb.HelloRequest{Name: name})
			if err != nil {
				log.Fatalf("call hello err: %v", err)
			}

			log.Println(r.Message)
		}
	}()

	for i := 0; i < 40000; i++ {
		r, err := fibclient.Fib(context.Background(), &pb.FibMsg{Num: int64(3)})
		if err != nil {
			log.Fatalf("call fib err: %v", err)
		}

		log.Printf("Fib(%d) = %v", i, r.Num)
	}

	wait.Wait()

}

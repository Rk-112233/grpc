package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"time"

	pb "example/demo"
	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.Dial("localhost:50051", grpc.WithInsecure())
	defer conn.Close()
	c := pb.NewDemoServiceClient(conn)

	// 1. Unary RPC
	unaryHello(c)

	// 2. Server Streaming
	serverStreamHello(c)

	// 3. Client Streaming
	clientStreamHello(c)

	// 4. BiDi Streaming
	bidiHello(c)
}

// Unary RPC
func unaryHello(c pb.DemoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	res, _ := c.UnaryHello(ctx, &pb.HelloRequest{Name: "Hemanth"})
	fmt.Println("Unary Response:", res.Message)
}

// Server Streaming RPC
func serverStreamHello(c pb.DemoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()
	stream, _ := c.ServerStreamHello(ctx, &pb.HelloRequest{Name: "Ranjith"})
	for {
		msg, err := stream.Recv()
		if err == io.EOF {
			break
		}
		fmt.Println("Server Stream:", msg.Message)
	}
}

// Client Streaming RPC
func clientStreamHello(c pb.DemoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	stream, _ := c.ClientStreamHello(ctx)

	names := []string{"Alice", "Bob", "Charlie"}
	for _, name := range names {
		stream.Send(&pb.HelloRequest{Name: name})
	}
	res, _ := stream.CloseAndRecv()
	fmt.Println("Client Stream Response:", res.Message)
}

// BiDi Streaming RPC
func bidiHello(c pb.DemoServiceClient) {
	ctx, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	stream, _ := c.BidiHello(ctx)

	done := make(chan struct{})

	// Receive messages
	go func() {
		for {
			res, err := stream.Recv()
			if err == io.EOF {
				break
			}
			fmt.Println("BiDi Response:", res.Message)
		}
		close(done)
	}()

	// Send messages
	names := []string{"X", "Y", "Z"}
	for _, name := range names {
		stream.Send(&pb.HelloRequest{Name: name})
		time.Sleep(time.Second)
	}
	stream.CloseSend()
	<-done
}

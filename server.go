package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	pb "example/demo" // import path where generated files are stored
	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedDemoServiceServer
}

// Unary RPC
func (s *server) UnaryHello(ctx context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "Hello " + req.Name}, nil
}

// Server Streaming RPC
func (s *server) ServerStreamHello(req *pb.HelloRequest, stream pb.DemoService_ServerStreamHelloServer) error {
	for i := 1; i <= 5; i++ {
		msg := fmt.Sprintf("Hello %s - Message %d", req.Name, i)
		if err := stream.Send(&pb.HelloResponse{Message: msg}); err != nil {
			return err
		}
		time.Sleep(1 * time.Second)
	}
	return nil
}

// Client Streaming RPC
func (s *server) ClientStreamHello(stream pb.DemoService_ClientStreamHelloServer) error {
	names := ""
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(&pb.HelloResponse{Message: "Hello " + names})
		}
		if err != nil {
			return err
		}
		names += req.Name + " "
	}
}

// BiDi Streaming RPC
func (s *server) BidiHello(stream pb.DemoService_BidiHelloServer) error {
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		reply := "Hello " + req.Name
		if err := stream.Send(&pb.HelloResponse{Message: reply}); err != nil {
			return err
		}
	}
}

func main() {
	lis, _ := net.Listen("tcp", ":50051")
	grpcServer := grpc.NewServer()
	pb.RegisterDemoServiceServer(grpcServer, &server{})
	fmt.Println("✅ gRPC Server started on :50051")
	log.Fatal(grpcServer.Serve(lis))
}

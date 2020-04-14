package main

import (
	pb "GRPCExample/ServerSideStreaming/ServerSideStreamingProto"
	"fmt"
	"log"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSimpleServiceServer
}

func (s *server) GetDivisors(req *pb.Request, stream pb.SimpleService_GetDivisorsServer) error {
	fmt.Println("Streaming started")
	num := req.Num
	fmt.Println("Number received:", num)
	var i int32
	for i = 1; i <= num; i++ {
		if num%i == 0 {
			fmt.Println("Sending:", i)
			if err := stream.Send(&pb.Response{Result: i}); err != nil {
				return err
			}
		}
	}
	return nil
}

func main() {
	lis, err := net.Listen("tcp", ":9093")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server started")
	g := grpc.NewServer()
	pb.RegisterSimpleServiceServer(g, &server{})
	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

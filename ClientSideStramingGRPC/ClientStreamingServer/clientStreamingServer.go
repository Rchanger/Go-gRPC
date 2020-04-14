package main

import (
	pb "GRPCExample/ClientSideStramingGRPC/ClientSideStramingProto"
	"fmt"
	"io"
	"log"
	"net"
	"time"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedSimpleServiceServer
}

func (s *server) FindMax(stream pb.SimpleService_FindMaxServer) error {
	var max int32
	startTime := time.Now()
	fmt.Println("startTime:", startTime)
	for {
		req, err := stream.Recv()
		if err == io.EOF {
			endTime := time.Now()
			fmt.Println("End time:", endTime)
			return stream.SendAndClose(&pb.Response{Result: max})
		}
		if err != nil {
			return err
		}
		fmt.Println("Received Num :", req.Num)
		if req.Num <= max {
			continue
		}
		max = req.Num
	}
}

func main() {
	lis, err := net.Listen("tcp", ":9092")
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

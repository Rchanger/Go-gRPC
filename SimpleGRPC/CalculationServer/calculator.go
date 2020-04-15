package main

import (
	"context"
	"fmt"
	"log"
	"net"

	pb "github.com/Rchanger/Go-gRPC/SimpleGRPC/ProtoDir"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalServiceServer
}

func (s *server) Add(ctx context.Context, requestData *pb.RequestData) (*pb.ResponseData, error) {
	firstNo := requestData.GetFirstNumber()
	secondNo := requestData.GetSecondNumber()
	addition := firstNo + secondNo
	result := pb.ResponseData{}
	result.Result = addition
	return &result, nil
}

func (s *server) Subtract(ctx context.Context, requestData *pb.RequestData) (*pb.ResponseData, error) {
	firstNo := requestData.GetFirstNumber()
	secondNo := requestData.GetSecondNumber()
	subtraction := firstNo - secondNo
	result := pb.ResponseData{}
	result.Result = subtraction
	return &result, nil
}

func main() {
	lis, err := net.Listen("tcp", ":9090")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("server started")
	s := grpc.NewServer()
	pb.RegisterCalServiceServer(s, &server{})
	if err := s.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

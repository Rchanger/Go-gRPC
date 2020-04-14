package main

import (
	pb "GRPCExample/BiDirectionGRPC/BiDirProto"
	"fmt"
	"io"
	"log"
	"net"

	"google.golang.org/grpc"
)

var count = 9

type server struct {
	pb.UnimplementedSimpleStreamServiceServer
}

func (s *server) Max(stream pb.SimpleStreamService_MaxServer) error {
	fmt.Println("Stream received")
	var max int32

	ctx := stream.Context()
	for {
		select {
		case <-ctx.Done():
			break
		default:
			// in, err := stream.Recv()
			// log.Println("Received value")
			// if err == io.EOF {
			// 	return nil
			// }
			// if err != nil {
			// 	return err
			// }
			// log.Println("Got " + in.Data)
		}
		req, err := stream.Recv()
		if err == io.EOF {
			log.Println("exit")
			return nil
		}
		if err != nil {
			log.Printf("receive error %v", err)
			continue
		}
		if req.Num <= max {
			continue
		}
		max = req.Num
		response := pb.Response{Result: max}
		if err := stream.Send(&response); err != nil {
			log.Printf("send error %v", err)
		}
		log.Printf("send new max=%d", max)
	}
}

func main() {
	lis, err := net.Listen("tcp", ":9091")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	fmt.Println("Server started")
	g := grpc.NewServer()
	pb.RegisterSimpleStreamServiceServer(g, &server{})
	if err := g.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}

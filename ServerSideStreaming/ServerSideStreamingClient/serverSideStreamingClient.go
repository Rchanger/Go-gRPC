package main

import (
	pb "GRPCExample/ServerSideStreaming/ServerSideStreamingProto"
	"context"
	"fmt"
	"io"
	"log"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9093", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("server started")
	defer conn.Close()
	c := pb.NewSimpleServiceClient(conn)
	req := pb.Request{Num: 50}
	stream, err := c.GetDivisors(context.Background(), &req)
	if err != nil {
		log.Fatalf("%v.GetDivisors(_) = _, %v", c, err)
	}
	for {
		divisor, err := stream.Recv()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatalf("%v.GetDivisors(_) = _, %v", c, err)
		}
		log.Println("divisor", divisor)
	}
}

package main

import (
	pb "GRPCExample/ClientSideStramingGRPC/ClientSideStramingProto"
	"context"
	"fmt"
	"log"
	"math/rand"
	"time"

	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().Unix())
	conn, err := grpc.Dial("localhost:9092", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("server started")
	defer conn.Close()
	c := pb.NewSimpleServiceClient(conn)
	stream, err := c.FindMax(context.Background())
	if err != nil {
		fmt.Println("err", err)
	}
	for i := 1; i <= 10; i++ {
		rnd := int32(rand.Intn(i))
		req := pb.Request{Num: rnd}
		fmt.Println("Sending no", rnd)
		if err := stream.Send(&req); err != nil {
			log.Fatalf("can not send %v", err)
		}
	}
	result, err := stream.CloseAndRecv()
	if err != nil {
		log.Fatalf("%v.CloseAndRecv() got error %v, want %v", stream, err, nil)
	}
	log.Printf("Max Number: %v", result)
}

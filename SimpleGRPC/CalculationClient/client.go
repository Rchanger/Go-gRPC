package main

import (
	pb "GRPCExample/SimpleGRPC/ProtoDir"
	"context"
	"log"
	"os"
	"strconv"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial("localhost:9090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	defer conn.Close()
	c := pb.NewCalServiceClient(conn)

	firstNo := 9
	secondNo := 5
	if len(os.Args) > 1 {
		firstNo, _ = strconv.Atoi(os.Args[1])
		secondNo, _ = strconv.Atoi(os.Args[2])
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	requestObj := pb.RequestData{}
	requestObj.FirstNumber = int64(firstNo)
	requestObj.SecondNumber = int64(secondNo)
	r, err := c.Add(ctx, &requestObj)
	if err != nil {
		log.Fatalf("could not get add response : %v", err)
	}
	log.Printf("Addition Result: %v", r.GetResult())
	r1, err1 := c.Subtract(ctx, &requestObj)
	if err1 != nil {
		log.Fatalf("could not get subtract response : %v", err1)
	}
	log.Printf("Subtraction Result: %v", r1.GetResult())
}

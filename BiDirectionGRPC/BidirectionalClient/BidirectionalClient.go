package main

import (
	"context"
	"fmt"
	"io"
	"log"
	"math/rand"
	"time"

	pb "GRPCExample/BiDirectionGRPC/BiDirProto"

	"google.golang.org/grpc"
)

func main() {
	rand.Seed(time.Now().Unix())

	conn, err := grpc.Dial("localhost:9091", grpc.WithInsecure())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}
	fmt.Println("server started")
	defer conn.Close()
	c := pb.NewSimpleStreamServiceClient(conn)
	stream, err := c.Max(context.Background())
	if err != nil {
		fmt.Println("err", err)
	}
	var max int32
	ctx := stream.Context()
	done := make(chan bool)

	// fmt.Println("msg", stream)
	// wait := make(chan struct{})
	// msg := pb.StramMessage{Data: "test"}
	// ctx := stream.Context()
	// count := 1
	/* go func() {
		for {
			log.Println("Sleeping...")
			time.Sleep(2 * time.Second)
			log.Println("Sending msg... ", count)
			count = count + 1
			stream.Send(&msg)
			if count > 15 {
				ctx.Done()
				break
			}
		}
	}()
	<-wait
	stream.CloseSend() */
	go func() {
		for i := 1; i <= 10; i++ {
			rnd := int32(rand.Intn(i))
			req := pb.Request{Num: rnd}
			if err := stream.Send(&req); err != nil {
				log.Fatalf("can not send %v", err)
			}
			log.Printf("%d sent", req.Num)
			time.Sleep(time.Millisecond * 200)
		}
		if err := stream.CloseSend(); err != nil {
			log.Println(err)
		}
	}()

	go func() {
		for {
			resp, err := stream.Recv()
			if err == io.EOF {
				close(done)
				return
			}
			if err != nil {
				log.Fatalf("can not receive %v", err)
			}
			max = resp.Result
			log.Printf("new max %d received", max)
		}
	}()

	go func() {
		<-ctx.Done()
		if err := ctx.Err(); err != nil {
			log.Println(err)
		}
		close(done)
	}()
	<-done
	log.Printf("finished with max=%d", max)

}

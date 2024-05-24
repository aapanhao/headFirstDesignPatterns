package main

import (
	pb "client/proto"
	"context"
	"log"
	"time"

	"google.golang.org/grpc"
)

func main() {
	conn, _ := grpc.NewClient("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	defer conn.Close()

	c := pb.NewGreeterClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, _ := c.SayHello(ctx, &pb.HelloRequest{Name: "world"})
	log.Println("Response:", r.Message)

}

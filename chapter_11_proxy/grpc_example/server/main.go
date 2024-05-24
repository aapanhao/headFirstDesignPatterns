package main

import (
	"context"
	"log"
	"net"
	pb "server/proto"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	return &pb.HelloReply{Message: "hello " + in.Name}, nil
}

func main() {

	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalln("listen err:", err)
	}

	if err := s.Serve(listen); err != nil {
		log.Fatalln("server err:", err)
	}

}

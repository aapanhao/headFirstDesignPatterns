package main

import (
	"log"
	"net"
	pb "server/proto"

	"google.golang.org/grpc"
)

func main() {

	machine := NewGumballMachine("beijing")
	machine.Refill(100)

	s := grpc.NewServer()

	pb.RegisterMonitorServer(s, machine)

	listen, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Println("err:", err)
	}

	if err := s.Serve(listen); err != nil {
		log.Println("err:", err)
	}
}

package main

import (
	"context"
	"log"

	pb "client/proto"

	"google.golang.org/grpc"
)

type GumballMonitor struct {
	GumballMachine pb.MonitorClient
}

func NewGumballMonitor() *GumballMonitor {
	g := &GumballMonitor{}
	conn, err := grpc.NewClient("localhost:50051", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalln("connect err:", err)
	}
	g.GumballMachine = pb.NewMonitorClient(conn)
	return g
}

func (g *GumballMonitor) report() {
	log.Println(g.GumballMachine.GetCount(context.Background(), &pb.Empty{}))
	log.Println(g.GumballMachine.GetLocation(context.Background(), &pb.Empty{}))
	log.Println(g.GumballMachine.GetState(context.Background(), &pb.Empty{}))
}

func main() {

	monitor := NewGumballMonitor()
	monitor.report()

}

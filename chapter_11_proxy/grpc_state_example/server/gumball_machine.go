package main

import (
	"context"
	"fmt"
	"log"
	"reflect"
	pb "server/proto"
)

type GumballMachine struct {
	SoldOutState State
	SoldState    State
	NoCoinState  State
	HasCoinState State
	WinnerState  State

	State    State
	Count    int
	Location string

	pb.UnimplementedMonitorServer
}

func (g *GumballMachine) GetCount(ctx context.Context, in *pb.Empty) (*pb.CountResponse, error) {
	return &pb.CountResponse{Count: int32(g.Count)}, nil
}

func (g *GumballMachine) GetLocation(ctx context.Context, in *pb.Empty) (*pb.LocatioinResponse, error) {
	return &pb.LocatioinResponse{Location: g.Location}, nil
}

func (g *GumballMachine) GetState(ctx context.Context, in *pb.Empty) (*pb.StateResponse, error) {
	return &pb.StateResponse{State: fmt.Sprintf("%v", reflect.TypeOf(g.State))}, nil
}

func (c *GumballMachine) String() string {
	return fmt.Sprintf("Client State: %v, Count: %v", reflect.TypeOf(c.State), c.Count)
}

func (c *GumballMachine) SetState(state State) {
	c.State = state
}

func (c *GumballMachine) Refill(count int) {
	c.Count = count
	c.State = c.NoCoinState
}

func (c *GumballMachine) ReleaseGumball() {
	log.Println("release a gumball")
	if c.Count > 0 {
		c.Count -= 1
	}
}

func (c *GumballMachine) InputCoin() {
	c.State.InputCoin()
}
func (c *GumballMachine) EjectCoin() {
	c.State.EjectCoin()
}
func (c *GumballMachine) TurnCrank() {
	c.State.TurnCrank()
	c.State.Dispense()
}

func NewGumballMachine(location string) *GumballMachine {
	client := GumballMachine{}
	client.SoldOutState = &SoldOutState{client: &client}
	client.SoldState = &SoldState{client: &client}
	client.NoCoinState = &NoCoinState{client: &client}
	client.HasCoinState = &HasCoinState{client: &client}
	client.WinnerState = &WinnerState{client: &client}
	client.Count = 0
	client.State = client.SoldOutState
	client.Location = location
	return &client
}

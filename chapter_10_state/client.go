package main

import (
	"fmt"
	"log"
	"reflect"
)

type Client struct {
	SoldOutState State
	SoldState    State
	NoCoinState  State
	HasCoinState State
	WinnerState  State

	State State
	Count int
}

func (c *Client) String() string {
	return fmt.Sprintf("Client State: %v, Count: %v", reflect.TypeOf(c.State), c.Count)
}

func (c *Client) SetState(state State) {
	c.State = state
}

func (c *Client) Refill(count int) {
	c.Count = count
	c.State = c.NoCoinState
}

func (c *Client) ReleaseGumball() {
	log.Println("release a gumball")
	if c.Count > 0 {
		c.Count -= 1
	}
}

func (c *Client) InputCoin() {
	c.State.InputCoin()
}
func (c *Client) EjectCoin() {
	c.State.EjectCoin()
}
func (c *Client) TurnCrank() {
	c.State.TurnCrank()
	c.State.Dispense()
}

func NewClient() *Client {
	client := Client{}
	client.SoldOutState = &SoldOutState{client: &client}
	client.SoldState = &SoldState{client: &client}
	client.NoCoinState = &NoCoinState{client: &client}
	client.HasCoinState = &HasCoinState{client: &client}
	client.WinnerState = &WinnerState{client: &client}
	client.Count = 0
	client.State = client.SoldOutState
	return &client
}

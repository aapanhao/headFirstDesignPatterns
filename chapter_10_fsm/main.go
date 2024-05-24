package main

import (
	"context"
	"log"

	"github.com/looplab/fsm"
)

var (
	soldOutState = "soldOutState"
	noCoinState  = "noCoinState"
	hasCoinState = "hasCoinState"
	soldState    = "soldState"
)

var (
	inputCoin    = "inputCoin"
	ejectCoin    = "ejectCoin"
	turnCrank    = "turnCrank"
	dispense     = "dispense"
	lastDispense = "lastDispense"
)

type GumballMachine struct {
	FSM   *fsm.FSM
	count int
}

func (g *GumballMachine) Refill(count int) {
	g.count = count
	g.FSM.SetState(noCoinState)
}

func (g *GumballMachine) inputCoin() {
	g.FSM.Event(context.Background(), inputCoin)

}

func (g *GumballMachine) GetState() {
	log.Println(g.FSM.Current())
}

func NewGumballMachine() *GumballMachine {
	g := GumballMachine{}
	g.count = 0
	g.FSM = fsm.NewFSM(
		soldOutState,
		fsm.Events{
			{Name: inputCoin, Src: []string{noCoinState}, Dst: hasCoinState},
			{Name: inputCoin, Src: []string{hasCoinState}, Dst: hasCoinState},
			{Name: inputCoin, Src: []string{soldOutState}, Dst: soldOutState},
			{Name: inputCoin, Src: []string{soldState}, Dst: soldState},
		},
		fsm.Callbacks{
			"before_inputCoin": func(ctx context.Context, e *fsm.Event) { beforeInputCoin(e) },
			"after_inputCoin":  func(ctx context.Context, e *fsm.Event) { afterInputCoin(e) },
		},
	)
	return &g
}
func beforeInputCoin(e *fsm.Event) {
	log.Printf("beforeInputCoin event: %s, src state:%s, dst state:%s\n", e.Event, e.Src, e.Dst)
}
func afterInputCoin(e *fsm.Event) {
	log.Printf("afterInputCoin event: %s, src state:%s, dst state:%s\n", e.Event, e.Src, e.Dst)
}

func main() {
	g := NewGumballMachine()
	g.GetState()

	g.Refill(5)
	g.GetState()

	log.Println("-----------start-------------")
	g.inputCoin()
	log.Println("------------ end ------------")

	log.Println("-----------start-------------")
	g.inputCoin()
	log.Println("------------ end ------------")

}

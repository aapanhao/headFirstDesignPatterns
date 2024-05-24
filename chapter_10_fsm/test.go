package main

// import (
// 	"context"
// 	"log"

// 	"github.com/looplab/fsm"
// )

// var (
// 	soldOutState = "soldOutState"
// 	noCoinState  = "noCoinState"
// 	hasCoinState = "hasCoinState"
// 	soldState    = "soldState"
// )

// var (
// 	inputCoin    = "inputCoin"
// 	ejectCoin    = "ejectCoin"
// 	turnCrank    = "turnCrank"
// 	dispense     = "dispense"
// 	lastDispense = "lastDispense"
// )

// type GumballMachine struct {
// 	FSM   *fsm.FSM
// 	count int
// }

// func NewGumballMachine() *GumballMachine {
// 	g := GumballMachine{}
// 	g.count = 0
// 	g.FSM = fsm.NewFSM(
// 		soldOutState,
// 		fsm.Events{
// 			{Name: inputCoin, Src: []string{noCoinState}, Dst: hasCoinState},
// 			// {Name: ejectCoin, Src: []string{hasCoinState}, Dst: noCoinState},
// 			// {Name: turnCrank, Src: []string{hasCoinState}, Dst: soldState},
// 			// {Name: dispense, Src: []string{soldState}, Dst: noCoinState},
// 			// {Name: lastDispense, Src: []string{soldState}, Dst: soldOutState},
// 		},
// 		fsm.Callbacks{
// 			"event_state": func(ctx context.Context, e *fsm.Event) {},
// 		},
// 	)
// 	return &g
// }

// func (g *GumballMachine) Refill(count int) {
// 	g.count = count
// 	g.FSM.SetState(noCoinState)
// }

// func (g *GumballMachine) inputCoin() {
// 	g.FSM.Event(context.Background(), inputCoin)

// }
// func (g *GumballMachine) ejectCoin() {
// 	g.FSM.Event(context.Background(), ejectCoin)
// }
// func (g *GumballMachine) turnCrank() {
// 	g.FSM.Event(context.Background(), turnCrank)
// 	if g.count > 1 {
// 		g.FSM.Event(context.Background(), dispense)
// 	} else {
// 		g.FSM.Event(context.Background(), lastDispense)
// 	}

// }
// func (g *GumballMachine) dispense() {

// }

// func main() {
// 	g := NewGumballMachine()
// 	log.Println(g.FSM.Current())

// 	g.Refill(5)
// 	log.Println(g.FSM.Current())

// 	g.inputCoin()

// }

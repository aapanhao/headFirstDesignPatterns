package main

import (
	"context"
	"log"

	"github.com/looplab/fsm"
)

func main() {
	fsm := fsm.NewFSM(
		"closed", fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "opening"},
			{Name: "close", Src: []string{"opening"}, Dst: "closed"},
		},
		fsm.Callbacks{},
	)
	log.Println(fsm.Current())

	fsm.Event(context.Background(), "open")
	log.Println(fsm.Current())
}

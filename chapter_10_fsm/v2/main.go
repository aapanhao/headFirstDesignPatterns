package main

import (
	"context"
	"log"

	"github.com/looplab/fsm"
)

type Door struct {
	To  string
	FSM *fsm.FSM
}

func NewDoor(to string) *Door {
	d := Door{To: to}
	d.FSM = fsm.NewFSM(
		"closed",
		fsm.Events{
			{Name: "open", Src: []string{"closed"}, Dst: "opening"},
			{Name: "close", Src: []string{"opening"}, Dst: "closed"},
			{Name: "haha", Src: []string{"closed"}, Dst: "closed"},
		},
		fsm.Callbacks{
			"before_open":  func(ctx context.Context, e *fsm.Event) { d.beforeOpen(e) },
			"before_close": func(ctx context.Context, e *fsm.Event) { d.beforeClose(e) },
			"before_event": func(ctx context.Context, e *fsm.Event) { d.beforeEvent(e) },

			"leave_opening": func(ctx context.Context, e *fsm.Event) { d.leaveOpening(e) },
			"leave_closed":  func(ctx context.Context, e *fsm.Event) { d.leaveClosed(e) },
			"leave_state":   func(ctx context.Context, e *fsm.Event) { d.leaveState(e) },

			"enter_opening": func(_ context.Context, e *fsm.Event) { d.enterOpening(e) },
			"enter_closed":  func(_ context.Context, e *fsm.Event) { d.enterClosed(e) },
			"enter_state":   func(_ context.Context, e *fsm.Event) { d.enterState(e) },

			"after_open":  func(ctx context.Context, e *fsm.Event) { d.afterOpen(e) },
			"after_close": func(ctx context.Context, e *fsm.Event) { d.afterClose(e) },
			"after_event": func(ctx context.Context, e *fsm.Event) { d.afterEvent(e) },
		},
	)
	return &d
}
func (d *Door) beforeOpen(e *fsm.Event) {
	log.Printf("beforeOpen event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}
func (d *Door) beforeClose(e *fsm.Event) {
	log.Printf("beforeClose event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}
func (d *Door) beforeEvent(e *fsm.Event) {
	log.Printf("beforeEvent event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}

func (d *Door) leaveOpening(e *fsm.Event) {
	log.Printf("leaveOpening event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}
func (d *Door) leaveClosed(e *fsm.Event) {
	log.Printf("leaveClosed event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}
func (d *Door) leaveState(e *fsm.Event) {
	log.Printf("leaveState event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}

func (d *Door) enterOpening(e *fsm.Event) {
	log.Printf("enterOpening event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}
func (d *Door) enterClosed(e *fsm.Event) {
	log.Printf("enterClosed event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}
func (d *Door) enterState(e *fsm.Event) {
	log.Printf("leaveState event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}
func (d *Door) afterOpen(e *fsm.Event) {
	log.Printf("afterOpen event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}
func (d *Door) afterClose(e *fsm.Event) {
	log.Printf("afterClose event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}
func (d *Door) afterEvent(e *fsm.Event) {
	log.Printf("afterEvent event:%s,  current:%s, dst:%s", e.Event, d.FSM.Current(), e.Dst)
}

func main() {

	d := NewDoor("heaven")

	// d.FSM.Event(context.Background(), "open")
	d.FSM.Event(context.Background(), "haha")

}

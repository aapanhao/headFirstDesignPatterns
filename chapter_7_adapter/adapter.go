package main

type AdapterTurkey struct {
	Turkey Turkey
}

func (a *AdapterTurkey) Quack() {
	a.Turkey.Gobble()
}
func (a *AdapterTurkey) Fly() {
	for i := 0; i < 5; i++ {
		a.Turkey.Fly()
	}
}

func NewAdapterTurkey(t Turkey) *AdapterTurkey {
	return &AdapterTurkey{t}
}

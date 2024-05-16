package main

type Waiter struct {
	MenuComponents MenuComponent
}

func (w *Waiter) Print() {
	w.MenuComponents.Print()
}

func NewWaiter(m MenuComponent) *Waiter {
	return &Waiter{m}
}

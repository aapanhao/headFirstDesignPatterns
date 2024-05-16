package main

import "log"

type A struct{}

func (*A) CreatePizza(pizzaType string) {
	log.Println("Get A")
}

func (p *A) OrderPizza(pizzaType string) {
	p.CreatePizza(pizzaType)
}

type B struct {
	A
}

func (*B) CreatePizza(pizzaType string) {
	log.Println("Get B")
}

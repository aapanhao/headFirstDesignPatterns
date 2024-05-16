package main

import "log"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}
type PizzaBase struct{}

func (*PizzaBase) Prepare() {
	log.Println("prepare pizza")
}
func (*PizzaBase) Bake() {
	log.Println("bake pizza")
}
func (*PizzaBase) Cut() {
	log.Println("cut pizza to six pieces")
}
func (*PizzaBase) Box() {
	log.Println("box pizza")
}

type NYCheesePizza struct {
	PizzaBase
}

func (*NYCheesePizza) Prepare() {
	log.Println("prepare cheese pizza")
}

type NYDurianPizza struct {
	PizzaBase
}

func (*NYDurianPizza) Prepare() {
	log.Println("prepare durian")
}

type ChicagoCheesePizza struct {
	PizzaBase
}

func (*ChicagoCheesePizza) Cut() {
	log.Println("cut pizza to four pieces")
}

type ChicagoDurianPizza struct {
	PizzaBase
}

func (*ChicagoDurianPizza) Prepare() {
	log.Println("prepare double durian")
}

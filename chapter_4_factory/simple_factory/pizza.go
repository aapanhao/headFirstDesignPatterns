package main

import "log"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type CheesePizza struct {
	Description string
}

func (*CheesePizza) Prepare() {
	log.Println("prepare cheese")
}

func (*CheesePizza) Bake() {
	log.Println("bake cheese pizza")
}

func (*CheesePizza) Cut() {
	log.Println("cut cheese pizza")
}

func (*CheesePizza) Box() {
	log.Println("box cheese pizza")
}

type DurianPizza struct {
	Description string
}

func (*DurianPizza) Prepare() {
	log.Println("prepare durian")
}

func (*DurianPizza) Bake() {
	log.Println("bake durian pizza")
}

func (*DurianPizza) Cut() {
	log.Println("cut durian pizza")
}

func (*DurianPizza) Box() {
	log.Println("box durian pizza")
}

type VeggiePizza struct {
	Description string
}

func (*VeggiePizza) Prepare() {
	log.Println("prepare veggie")
}

func (*VeggiePizza) Bake() {
	log.Println("bake veggie pizza")
}

func (*VeggiePizza) Cut() {
	log.Println("cut veggie pizza")
}

func (*VeggiePizza) Box() {
	log.Println("box veggie pizza")
}

type DefaultPizza struct {
	Description string
}

func (*DefaultPizza) Prepare() {
	log.Println("not prepare")
}

func (*DefaultPizza) Bake() {
	log.Println("bake default pizza")
}

func (*DefaultPizza) Cut() {
	log.Println("cut default pizza")
}

func (*DefaultPizza) Box() {
	log.Println("box default pizza")
}

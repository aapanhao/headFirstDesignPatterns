package main

import "log"

type Turkey interface {
	Gobble()
	Fly()
}

type WildTurkey struct{}

func (*WildTurkey) Gobble() {
	log.Println("wild turkey gobble")
}

func (*WildTurkey) Fly() {
	log.Println("wild turkey fly short distance")
}

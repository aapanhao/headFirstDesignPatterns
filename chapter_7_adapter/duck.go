package main

import "log"

type Duck interface {
	Quack()
	Fly()
}

type MallardDuck struct{}

func (*MallardDuck) Quack() {
	log.Println("mallard duck quack")
}

func (*MallardDuck) Fly() {
	log.Println("mallard duck fly")
}

package main

import "log"

type Quack interface {
	quack()
}

type RedDuck struct{}

func (r *RedDuck) quack() {
	log.Println("red duck quack")
}

type GreedDuck struct{}

func (r *GreedDuck) quack() {
	log.Println("green duck quack")
}

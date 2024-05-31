package main

import "log"

type RedHeadDuck struct {
	observable Observable
}

func (r *RedHeadDuck) quack() {
	log.Println("red head duck quack")
	r.notify()
}

func NewRedHeadDuck() *RedHeadDuck {
	r := &RedHeadDuck{}
	r.observable = Observable{duck: r}
	return r
}

func (r *RedHeadDuck) register(observer Observer) {
	r.observable.register(observer)
}
func (r *RedHeadDuck) notify() {
	r.observable.notify()
}

type GreenDuck struct{}

func (*GreenDuck) quack() {
	log.Println("green head duck quack")
}

type RubberDuck struct{}

func (*RubberDuck) quack() {
	log.Println("kuwk")
}

type GooseAdapter struct {
	goose Goose
}

func (g *GooseAdapter) quack() {
	g.goose.honk()
}

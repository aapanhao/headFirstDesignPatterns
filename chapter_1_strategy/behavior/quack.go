package behavior

import "log"

type QuackBehavior interface {
	Quack()
}

type BigQuack struct{}

func (*BigQuack) Quack() {
	log.Println("i quack loudly")
}

type SlientQuack struct{}

func (*SlientQuack) Quack() {
	log.Println("i can't quack")
}

type NormalQuack struct{}

func (*NormalQuack) Quack() {
	log.Println("im quack")
}

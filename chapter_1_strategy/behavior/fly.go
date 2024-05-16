package behavior

import "log"

type FlyBehavior interface {
	Fly()
}

type RocketFly struct{}

func (*RocketFly) Fly() {
	log.Println("im fly with rocket")
}

type NoFly struct{}

func (*NoFly) Fly() {
	log.Println("i can't fly")
}

type NormalFly struct{}

func (*NormalFly) Fly() {
	log.Println("im flying")
}

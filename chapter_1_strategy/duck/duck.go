package duck

import "github.com/aapanhao/headFirstDesignPatterns/chapter_1_strategy_pattern/behavior"

type Duck struct {
	FlyBehavior   behavior.FlyBehavior
	QuackBehavior behavior.QuackBehavior
}

func (d *Duck) PerformFly() {
	d.FlyBehavior.Fly()
}

func (d *Duck) PerformQuack() {
	d.QuackBehavior.Quack()
}

type RedDuck struct {
	Duck
}

type RocketDuck struct {
	Duck
}

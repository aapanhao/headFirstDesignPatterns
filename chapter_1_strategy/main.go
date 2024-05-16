package main

import (
	"github.com/aapanhao/headFirstDesignPatterns/chapter_1_strategy_pattern/behavior"
	"github.com/aapanhao/headFirstDesignPatterns/chapter_1_strategy_pattern/duck"
)

func main() {
	r := duck.RedDuck{
		Duck: duck.Duck{
			FlyBehavior:   &behavior.RocketFly{},
			QuackBehavior: &behavior.NormalQuack{},
		},
	}
	r.PerformFly()
	r.PerformQuack()
}

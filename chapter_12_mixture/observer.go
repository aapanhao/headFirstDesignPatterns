package main

import (
	"log"
	"reflect"
)

type Observer interface {
	update(QuackObservable)
}

type DuckObserver struct{}

func (*DuckObserver) update(o QuackObservable) {
	log.Println("i get notify", reflect.TypeOf(o), "quack")
}

package main

import (
	"log"
	"reflect"
)

type Waiter struct {
	m MenuIterator
}

func NewWaiter(m MenuIterator) *Waiter {
	return &Waiter{m}
}

func (w *Waiter) Server() {

	log.Printf("-----------%v menu-------------", reflect.TypeOf(w.m))
	i := w.m.GetMenuIterator()
	PrintMenu(i)
	log.Println("-----------------------------------")
}

func PrintMenu(i Iterator) {
	for i.HasNext() {
		menuItem := i.Next()
		log.Println(menuItem.Description, menuItem.Price, menuItem.Vegetarian)
	}
}

package main

import "log"

type PizzaStore interface {
	CreatePizza(string) Pizza
}

func OrderPizza(pizzaStore PizzaStore, pizzaType string) {
	pizza := pizzaStore.CreatePizza(pizzaType)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
}

type PizzaStoreBase struct{}

func (*PizzaStoreBase) CreatePizza(pizzaType string) Pizza {
	log.Println("Get PizzaBase")
	return &PizzaBase{}
}

type NYPizzaStore struct {
	PizzaStore
}

func (*NYPizzaStore) CreatePizza(pizzaType string) Pizza {
	log.Println("Get NYPizza")

	if pizzaType == "cheese" {
		return &NYCheesePizza{}
	}
	if pizzaType == "durain" {
		return &NYDurianPizza{}
	}
	return &PizzaBase{}
}

type ChicagoPizzaStore struct {
	PizzaStore
}

func (*ChicagoPizzaStore) CreatePizza(pizzaType string) Pizza {
	if pizzaType == "cheese" {
		return &ChicagoCheesePizza{}
	}
	if pizzaType == "durain" {
		return &ChicagoDurianPizza{}
	}
	return &PizzaBase{}
}

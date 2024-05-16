package main

type SimplePizzaFactory struct {
}

func (*SimplePizzaFactory) CreatePizza(pizzaType string) Pizza {
	if pizzaType == "cheese" {
		return &CheesePizza{}
	}
	if pizzaType == "durian" {
		return &DurianPizza{}
	}
	if pizzaType == "veggie" {
		return &VeggiePizza{}
	}
	return &DefaultPizza{}
}

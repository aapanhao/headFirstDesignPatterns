package main

type PizzaStore struct {
	PizzaFactory SimplePizzaFactory
}

func NewPizzaStore(pizzaFactory SimplePizzaFactory) PizzaStore {
	return PizzaStore{PizzaFactory: pizzaFactory}
}

func (p *PizzaStore) OrderPizza(pizzaType string) Pizza {
	pizza := p.PizzaFactory.CreatePizza(pizzaType)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
	return pizza
}

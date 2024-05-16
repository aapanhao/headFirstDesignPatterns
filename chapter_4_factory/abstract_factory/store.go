package main

type PizzaStore interface {
	CreatePizza(string) Pizza
}

type PizzaStoreBase struct{}

func (p *PizzaStoreBase) CreatePizza(pizzaType string) Pizza {
	factory := PizzaIngredientFactoryBase{}
	return NewBasePizza(&factory)
}

func OrderPizza(pizzaType string, pizzaStore PizzaStore) {
	pizza := pizzaStore.CreatePizza(pizzaType)
	pizza.Prepare()
	pizza.Bake()
	pizza.Cut()
	pizza.Box()
}

type NYPizzaStore struct {
	PizzaStoreBase
}

func (n *NYPizzaStore) CreatePizza(pizzaType string) Pizza {
	factory := &NYPizzaIngredientFactory{}
	if pizzaType == "cheese" {
		return NewCheesePizza(factory)
	}
	if pizzaType == "veggie" {
		return NewVeggiePizza(factory)
	}
	return &PizzaBase{}
}

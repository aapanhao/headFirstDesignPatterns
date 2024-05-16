package main

import "log"

type Pizza interface {
	Prepare()
	Bake()
	Cut()
	Box()
}

type PizzaBase struct {
	Factory PizzaIngredientFactory
}

func NewBasePizza(factory PizzaIngredientFactory) Pizza {
	return &PizzaBase{Factory: factory}
}

func (c *PizzaBase) Prepare() {
	sauce := c.Factory.CreateSauce()
	cheese := c.Factory.CreateCheese()
	claim := c.Factory.CreateClaim()
	log.Println("prepare", sauce, cheese, claim)
}
func (c *PizzaBase) Bake() {
	log.Println("bake base pizza")
}

func (c *PizzaBase) Cut() {
	log.Println("cut base pizza")
}

func (c *PizzaBase) Box() {
	log.Println("box base pizza")
}

type CheesePizza struct {
	PizzaBase
}

func NewCheesePizza(factory PizzaIngredientFactory) *CheesePizza {
	return &CheesePizza{PizzaBase{Factory: factory}}
}

type VeggiePizza struct {
	PizzaBase
}

func NewVeggiePizza(factory PizzaIngredientFactory) *VeggiePizza {
	return &VeggiePizza{PizzaBase{Factory: factory}}
}

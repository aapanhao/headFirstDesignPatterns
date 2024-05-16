package main

type PizzaIngredientFactory interface {
	CreateSauce() Sauce
	CreateCheese() Cheese
	CreateClaim() Claim
}

type PizzaIngredientFactoryBase struct {
}

func (n *PizzaIngredientFactoryBase) CreateSauce() Sauce {
	return BaseSauce{"base sauce"}
}
func (n *PizzaIngredientFactoryBase) CreateCheese() Cheese {
	return BaseCheese{"base cheese"}
}
func (n *PizzaIngredientFactoryBase) CreateClaim() Claim {
	return BaseClaim{"base claim"}
}

type NYPizzaIngredientFactory struct {
	PizzaIngredientFactoryBase
}

func (n *NYPizzaIngredientFactory) CreateSauce() Sauce {
	return ThinSauce{"thin sauce"}
}
func (n *NYPizzaIngredientFactory) CreateCheese() Cheese {
	return ThinCheese{"thin cheese"}
}
func (n *NYPizzaIngredientFactory) CreateClaim() Claim {
	return ThinClaim{"thin claim"}
}

type ChicagoPizzaIngredientFactory struct{}

func (n *ChicagoPizzaIngredientFactory) CreateSauce() Sauce {
	return ThickSauce{"thick sauce"}
}
func (n *ChicagoPizzaIngredientFactory) CreateCheese() Cheese {
	return ThickCheese{"thick cheese"}
}
func (n *ChicagoPizzaIngredientFactory) CreateClaim() Claim {
	return ThickClaim{"thick claim"}
}

type Sauce interface{}
type BaseSauce struct {
	Description string
}
type ThinSauce struct {
	Description string
}
type ThickSauce struct {
	Description string
}

type Cheese interface{}
type BaseCheese struct{ Description string }
type ThinCheese struct{ Description string }
type ThickCheese struct{ Description string }

type Claim interface{}
type BaseClaim struct{ Description string }
type ThinClaim struct{ Description string }
type ThickClaim struct{ Description string }

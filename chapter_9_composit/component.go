package main

type MenuComponent interface {
	GetName() string
	GetDescription() string
	GetPrice() float64
	IsVegetarain() bool
	Print()
	Add(...MenuComponent)
	Remove(MenuComponent)
	GetChild(int) MenuComponent
}

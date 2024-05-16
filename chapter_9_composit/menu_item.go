package main

import "log"

type MenuItem struct {
	MenuComponent
	description string
	name        string
	price       float64
	vegetarain  bool
}

func (m *MenuItem) GetDescription() string {
	return m.description
}
func (m *MenuItem) GetName() string {
	return m.name
}
func (m *MenuItem) GetPrice() float64 {
	return m.price
}
func (m *MenuItem) IsVegetarain() bool {
	return m.vegetarain
}

func (m *MenuItem) Print() {
	log.Printf("	MenuItem Name:%s, IsVegetarain:%v, Description:%s, Price:%.2f\n",
		m.GetName(), m.IsVegetarain(), m.GetDescription(), m.GetPrice())
}

func NewMenuItem(name, description string, price float64, vegetarain bool) *MenuItem {
	return &MenuItem{
		description: description,
		name:        name,
		price:       price,
		vegetarain:  vegetarain,
	}
}

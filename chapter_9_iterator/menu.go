package main

type MenuItem struct {
	Description string
	Price       float64
	Vegetarian  bool
}

type MenuIterator interface {
	GetMenuIterator() Iterator
}

type Dinner struct {
	MenuItems []MenuItem
}

func (d *Dinner) AddMenu(m MenuItem) {
	d.MenuItems = append(d.MenuItems, m)
}

func NewDinner() *Dinner {
	d := Dinner{}
	d.AddMenu(MenuItem{"pizza", 18, false})
	d.AddMenu(MenuItem{"salad", 12, true})
	d.AddMenu(MenuItem{"noddle", 16, false})
	return &d
}

func (d *Dinner) GetMenu() []MenuItem {
	return d.MenuItems
}

func (d *Dinner) GetMenuIterator() Iterator {
	return &DinnerIterator{d.MenuItems, 0}
}

type Pancake struct {
	MenuItems map[string]MenuItem
}

func (p *Pancake) AddMenu(m MenuItem) {
	p.MenuItems[m.Description] = m
}

func NewPancake() *Pancake {
	p := Pancake{MenuItems: make(map[string]MenuItem)}
	p.AddMenu(MenuItem{"egg cake", 6, true})
	p.AddMenu(MenuItem{"meat cake", 8, false})
	p.AddMenu(MenuItem{"vegetable cake", 6, false})
	return &p
}

func (p *Pancake) GetMenu() map[string]MenuItem {
	return p.MenuItems
}

func (d *Pancake) GetMenuIterator() Iterator {
	return NewPancakeIterator(d.MenuItems)
}

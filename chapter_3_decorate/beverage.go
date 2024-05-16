package main

type Beverage interface {
	GetDiscription() string
	Cost() float64
}

type Tea struct {
	Description string
}

func (t *Tea) GetDiscription() string {
	return t.Description
}

func (t *Tea) Cost() float64 {
	return 10
}

func NewTea() Beverage {
	return &Tea{Description: "tea"}
}

type RoastCoffee struct {
	Description string
}

func (r *RoastCoffee) GetDiscription() string {
	return r.Description
}

func (r *RoastCoffee) Cost() float64 {
	return 12.5
}

func NewRoastCoffee() Beverage {
	return &RoastCoffee{"roast coffee"}
}

type Decaf struct {
	Description string
}

func (d *Decaf) GetDiscription() string {
	return d.Description
}

func (d *Decaf) Cost() float64 {
	return 12.5
}

func NewDecaf() Beverage {
	return &Decaf{"decaf"}
}

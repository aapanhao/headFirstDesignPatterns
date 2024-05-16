package main

type Milk struct {
	Beverage    Beverage
	Description string
}

func (m *Milk) GetDiscription() string {
	return m.Beverage.GetDiscription() + m.Description
}

func (m *Milk) Cost() float64 {
	return m.Beverage.Cost() + 4.2
}

func AddMilk(b Beverage) Beverage {
	return &Milk{b, " add milk"}
}

type Mocha struct {
	Beverage    Beverage
	Description string
}

func (m *Mocha) GetDiscription() string {
	return m.Beverage.GetDiscription() + m.Description
}

func (m *Mocha) Cost() float64 {
	return m.Beverage.Cost() + 3.6
}

func AddMocha(b Beverage) Beverage {
	return &Mocha{b, " add mocha"}
}

type Sugar struct {
	Beverage    Beverage
	Description string
}

func (s *Sugar) GetDiscription() string {
	return s.Beverage.GetDiscription() + s.Description
}

func (s *Sugar) Cost() float64 {
	return s.Beverage.Cost() + 0.2
}

func AddSugar(b Beverage) Beverage {
	return &Sugar{b, " add sugar"}
}

package main

type Iterator interface {
	HasNext() bool
	Next() MenuItem
}

type DinnerIterator struct {
	MenuItems []MenuItem
	position  int
}

func (d *DinnerIterator) HasNext() bool {
	return d.position < len(d.MenuItems)
}

func (d *DinnerIterator) Next() MenuItem {
	m := d.MenuItems[d.position]
	d.position += 1
	return m
}

type PancakeIterator struct {
	MenuItems        map[string]MenuItem
	MenuDescriptions []string
	position         int
}

func NewPancakeIterator(m map[string]MenuItem) *PancakeIterator {
	menuDescriptions := []string{}
	for name := range m {
		menuDescriptions = append(menuDescriptions, name)
	}
	return &PancakeIterator{m, menuDescriptions, 0}
}

func (p *PancakeIterator) HasNext() bool {
	return p.position < len(p.MenuItems)
}

func (p *PancakeIterator) Next() MenuItem {
	description := p.MenuDescriptions[p.position]
	m := p.MenuItems[description]
	p.position += 1
	return m
}

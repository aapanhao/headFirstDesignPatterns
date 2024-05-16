package main

import "log"

type Menu struct {
	MenuComponent
	description    string
	name           string
	MenuComponents []MenuComponent
}

func (m *Menu) GetDescription() string {
	return m.description
}

func (m *Menu) GetName() string {
	return m.name
}

func (m *Menu) Add(menus ...MenuComponent) {
	m.MenuComponents = append(m.MenuComponents, menus...)
}

func (m *Menu) Remove(menu MenuComponent) {

}

func (m *Menu) GetChild(i int) MenuComponent {
	return m.MenuComponents[i]
}

func (m *Menu) Print() {
	log.Printf("Menu Name:%s, Description:%s\n", m.GetName(), m.GetDescription())
	for _, m := range m.MenuComponents {
		m.Print()
	}
}

func NewMenu(name, description string) *Menu {
	return &Menu{description: description, name: name}
}

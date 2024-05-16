package main

import "log"

type Light struct {
	Name string
}

func (l *Light) On() {
	log.Printf("%s light is on\n", l.Name)
}

func (l *Light) Off() {
	log.Printf("%s light is off\n", l.Name)
}

type TV struct {
	Name string
}

func (t *TV) Open() {
	log.Printf("%s tv is open\n", t.Name)
}

func (t *TV) Close() {
	log.Printf("%s tv is close\n", t.Name)
}

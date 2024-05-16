package main

import "log"

type Dvd struct{}

func (*Dvd) On() {
	log.Println("dvd on")
}
func (*Dvd) Off() {
	log.Println("dvd off")
}

type Cd struct{}

func (*Cd) On() {
	log.Println("cd on")
}
func (*Cd) Off() {
	log.Println("cd off")
}

type Light struct{}

func (*Light) Open() {
	log.Println("light open")
}
func (*Light) Close() {
	log.Println("light close")
}
func (*Light) SetLight(luminance int) {
	log.Println("light is set", luminance)
}

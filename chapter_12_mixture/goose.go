package main

import "log"

type Goose struct{}

func (*Goose) honk() {
	log.Println("goose honk")
}

package main

import "log"

func main() {
	s1 := NewSingleton()
	s2 := NewSingleton()
	log.Println(s1 == s2)
}

package main

import "log"

func main() {
	tea := NewTea()
	b := AddSugar(AddMilk(AddMilk(tea)))
	log.Println("description:", b.GetDiscription(), "cost:", b.Cost())
}

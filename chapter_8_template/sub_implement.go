package main

import "log"

type Tea struct {
	BaseBeverage
}

func (t *Tea) Step2() {
	log.Println("Tea step 2")
}

func (t *Tea) Step4() {
	log.Println("Tea step 4")
}

type Coffee struct {
	BaseBeverage
}

func (t *Coffee) Step2() {
	log.Println("Coffee step 2")
}

func (t *Coffee) Step4() {
	log.Println("Coffee step 4")
}

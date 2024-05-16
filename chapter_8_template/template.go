package main

import "log"

type Beverge interface {
	Step1()
	Step2()
	Step3()
	Step4()
}

type BaseBeverage struct{}

func (b *BaseBeverage) Step1() {
	log.Println("Base Step1")
}
func (b *BaseBeverage) Step2() {
	log.Println("Base Step2")
}
func (b *BaseBeverage) Step3() {
	log.Println("Base Step3")
}
func (b *BaseBeverage) Step4() {
	log.Println("Base Step4")
}

func Prepare(b Beverge) {
	b.Step1()
	b.Step2()
	b.Step3()
	b.Step4()
}

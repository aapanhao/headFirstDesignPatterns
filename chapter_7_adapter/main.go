package main

func TestDuck(duck Duck) {
	duck.Quack()
	duck.Fly()
}

func main() {
	d := MallardDuck{}
	TestDuck(&d)

	t := WildTurkey{}
	adapterTurkey := NewAdapterTurkey(&t)
	TestDuck(adapterTurkey)
}

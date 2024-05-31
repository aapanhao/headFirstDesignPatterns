package main

func simulate(duck Quackable) {
	duck.quack()
}

func main() {
	r := NewRedHeadDuck()
	r.register(&DuckObserver{})
	simulate(r)

}

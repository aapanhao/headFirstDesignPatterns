package main

type FlockDuck struct {
	ducks []Quackable
}

func (f *FlockDuck) add(duck Quackable) {
	f.ducks = append(f.ducks, duck)
}

func (f *FlockDuck) quack() {
	for _, duck := range f.ducks {
		duck.quack()
	}
}

package main

type QuackObservable interface {
	register(Observer)
	notify()
}

type Observable struct {
	observers []Observer
	duck      QuackObservable
}

func (o *Observable) register(observer Observer) {
	o.observers = append(o.observers, observer)
}

func (o *Observable) notify() {
	for _, observer := range o.observers {
		observer.update(o.duck)
	}
}

package main

func main() {
	simpleFactory := SimplePizzaFactory{}
	pizzaStore := NewPizzaStore(simpleFactory)
	pizzaStore.OrderPizza("veggie")
}

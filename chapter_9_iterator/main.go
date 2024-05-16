package main

func main() {
	dinner := NewDinner()
	waiter := NewWaiter(dinner)
	waiter.Server()
}

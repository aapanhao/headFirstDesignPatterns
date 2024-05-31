package main

import "log"

func main() {
	NewCounter(&RedDuck{}).quack()
	NewCounter(&RedDuck{}).quack()
	NewCounter(&GreedDuck{}).quack()
	NewCounter(&GreedDuck{}).quack()
	NewCounter(&GreedDuck{}).quack()
	log.Println("quack times:", quackCount)
}

package main

import "log"

func main() {
	client := NewClient()
	log.Println(client)

	client.Refill(5)
	log.Println(client)

	client.InputCoin()
	client.TurnCrank()
	log.Println(client)

	client.InputCoin()
	client.TurnCrank()
	log.Println(client)

	client.InputCoin()
	client.TurnCrank()
	log.Println(client)
}

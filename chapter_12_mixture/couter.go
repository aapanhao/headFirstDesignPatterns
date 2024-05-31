package main

var quackCount int

type Counter struct {
	duck Quackable
}

func (c *Counter) quack() {
	quackCount++
	c.duck.quack()
}

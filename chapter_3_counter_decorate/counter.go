package main

var (
	quackCount int
)

type Counter struct {
	Quack Quack
}

func NewCounter(quack Quack) *Counter {
	return &Counter{Quack: quack}
}

func (c *Counter) quack() {
	quackCount++
	c.Quack.quack()
}

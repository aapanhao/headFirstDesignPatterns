package main

import (
	"fmt"
	"reflect"
)

type Client struct {
	OnCommands  [7]Command
	OffCommands [7]Command
	preCommand  Command
}

func NewClient() *Client {
	c := Client{}
	for i := 0; i < 7; i++ {
		c.OnCommands[i] = &NoCommand{}
		c.OffCommands[i] = &NoCommand{}
	}
	return &c
}

func (c *Client) String() string {
	result := ""
	result += "\n-------Remote Controller-------\n"
	for i := 0; i < 7; i++ {
		onCommand := reflect.TypeOf(c.OnCommands[i])
		offCommand := reflect.TypeOf(c.OffCommands[i])
		result += fmt.Sprintf("[slot %d] %v %v\n", i, onCommand, offCommand)
	}
	result += "\n"
	return result
}

func (c *Client) SetCommand(slot int, onCommand Command, offCommand Command) {
	c.OnCommands[slot] = onCommand
	c.OffCommands[slot] = offCommand
}

func (c *Client) OnButtonWasPressed(slot int) {
	c.preCommand = c.OnCommands[slot]
	c.OnCommands[slot].Execute()
}

func (c *Client) OffButtonWasPressed(slot int) {
	c.preCommand = c.OffCommands[slot]
	c.OffCommands[slot].Execute()
}

func (c *Client) UndoButtonWasPressed() {
	c.preCommand.Undo()
}

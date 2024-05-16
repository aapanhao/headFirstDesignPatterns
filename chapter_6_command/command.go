package main

import "log"

type Command interface {
	Execute()
	Undo()
}

type NoCommand struct{}

func (*NoCommand) Execute() {
	log.Println("no command")
}
func (*NoCommand) Undo() {
	log.Println("no command")
}

type LightOnCommand struct {
	Light *Light
}

func (l *LightOnCommand) Execute() {
	l.Light.On()
}
func (l *LightOnCommand) Undo() {
	l.Light.Off()
}

func NewLightOnCommand(l *Light) *LightOnCommand {
	return &LightOnCommand{l}
}

type LightOffCommand struct {
	Light *Light
}

func (l *LightOffCommand) Execute() {
	l.Light.Off()
}
func (l *LightOffCommand) Undo() {
	l.Light.On()
}

func NewLightOffCommand(l *Light) *LightOffCommand {
	return &LightOffCommand{l}
}

type TVOnCommand struct {
	TV *TV
}

func (t *TVOnCommand) Execute() {
	t.TV.Open()
}
func (t *TVOnCommand) Undo() {
	t.TV.Close()
}
func NewTVOnCommand(t *TV) *TVOnCommand {
	return &TVOnCommand{t}
}

type TVOffCommand struct {
	TV *TV
}

func (t *TVOffCommand) Execute() {
	t.TV.Close()
}
func (t *TVOffCommand) Undo() {
	t.TV.Open()
}

func NewTVOffCommand(t *TV) *TVOffCommand {
	return &TVOffCommand{t}
}

type MarcoCommand struct {
	Commands []Command
}

func (m *MarcoCommand) Execute() {
	for _, c := range m.Commands {
		c.Execute()
	}

}
func (m *MarcoCommand) Undo() {
	for _, c := range m.Commands {
		c.Undo()
	}
}

func NewMarcoCommand(commands ...Command) *MarcoCommand {
	return &MarcoCommand{commands}
}

type MarcoOffCommand struct {
	Commands []Command
}

func (m *MarcoOffCommand) Execute() {
	for _, c := range m.Commands {
		c.Execute()
	}

}
func (m *MarcoOffCommand) Undo() {
	for _, c := range m.Commands {
		c.Undo()
	}
}

func NewMarcoOffCommand(commands ...Command) *MarcoOffCommand {
	return &MarcoOffCommand{commands}
}

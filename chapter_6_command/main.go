package main

import "log"

func main() {

	light := Light{"room"}
	lightOnCommand := NewLightOnCommand(&light)
	lightOffCommand := NewLightOffCommand(&light)

	tv := TV{"living room"}
	tvOnCommand := NewTVOnCommand(&tv)
	tvOffCommand := NewTVOffCommand(&tv)

	macroOnCommand := NewMarcoCommand(lightOnCommand, tvOnCommand)
	macroOffCommand := NewMarcoOffCommand(lightOffCommand, tvOffCommand)

	client := NewClient()
	client.SetCommand(0, lightOnCommand, lightOffCommand)
	client.SetCommand(1, tvOnCommand, tvOffCommand)
	client.SetCommand(2, macroOnCommand, macroOffCommand)

	log.Println(client)

	client.OnButtonWasPressed(2)
	// client.OffButtonWasPressed(2)
	// client.UndoButtonWasPressed()

}

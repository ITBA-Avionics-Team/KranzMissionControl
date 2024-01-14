package main

import (
	"kranz/communication-module/broadcast"
	"kranz/communication-module/controller"
	"kranz/communication-module/model"
	"kranz/communication-module/serial"
)

func main() {
	systemStatusBroadcast := broadcast.NewBroadcast[model.SystemStatus]()
	serialPort := "/dev/ttys008"
	go serial.ListenForMessages(serialPort, systemStatusBroadcast)
	controller.NewCommunicationModuleRouter(*systemStatusBroadcast).Run(":8080")
}

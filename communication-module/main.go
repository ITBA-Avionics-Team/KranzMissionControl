package main

import (
	"kranz/communication-module/broadcast"
	"kranz/communication-module/controller"
	"kranz/communication-module/model"
	"kranz/communication-module/serial"
	"log"
)

func main() {
	systemStatusBroadcast := broadcast.NewBroadcast[model.SystemStatus]()
	serialPort := "/dev/ttys008"
	port, mutex, err := serial.OpenSerialConnection(serialPort)
	if err != nil {
		log.Fatalf("failed to opn serial port: %v", err)
	}
	go serial.ListenForMessages(port, &mutex, systemStatusBroadcast)
	controller.NewCommunicationModuleRouter(*systemStatusBroadcast, port, &mutex).Run(":8080")
	defer port.Close()
}

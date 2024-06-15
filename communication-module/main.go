package main

import (
	"fmt"
	"kranz/communication-module/broadcast"
	"kranz/communication-module/controller"
	"kranz/communication-module/model"
	"kranz/communication-module/serial"
	"log"
	"os"
)

func main() {
	var serialPortName string
	args := os.Args
    if len(args) > 1 {
        serialPortName = args[1]
    } else {
        fmt.Println("No serial port specified. Exiting...")
				os.Exit(1)
    }
	
	systemStatusBroadcast := broadcast.NewBroadcast[model.SystemStatus]()
	serialPort, mutex, err := serial.OpenSerialConnection(serialPortName)
	serialPort.SetReadTimeout(1000000);
	if err != nil {
		log.Fatalf("failed to opn serial port: %v", err)
	}
	go serial.ListenForMessages(serialPort, &mutex, systemStatusBroadcast)
	controller.NewCommunicationModuleRouter(*systemStatusBroadcast, serialPort, &mutex).Run(":8080")
	defer serialPort.Close()
}

package main

import (
	"fmt"
	"kranz/communication-module/broadcast"
	"kranz/communication-module/model"
	"kranz/communication-module/serial"
	"log"
	"os"
	"os/signal"
	"syscall"
)

func main() {
	fmt.Println("=== INICIANDO LECTURA DE TELEMETRÍA DESDE XBEE ===")

	// Definir el puerto serial (COM10 en Windows)
	serialPort := "COM10"
	fmt.Printf("Intentando conectar al puerto: %s\n", serialPort)

	// Abrir conexión serial
	port, mutex, err := serial.OpenSerialConnection(serialPort)
	if err != nil {
		log.Fatalf("Error abriendo puerto serial: %v", err)
	}
	defer port.Close()

	fmt.Println("Conexión serial establecida exitosamente")

	// Crear broadcast para enviar los mensajes parseados
	systemStatusBroadcast := broadcast.NewBroadcast[model.SystemStatus]()

	// Suscribirse al broadcast para recibir los mensajes parseados
	subscriptionId, subscription := systemStatusBroadcast.Subscribe()
	defer systemStatusBroadcast.Unubscribe(subscriptionId)

	// Iniciar la escucha de mensajes en una goroutine
	go serial.ListenForMessages(port, &mutex, systemStatusBroadcast)

	// Goroutine para imprimir los mensajes recibidos
	go func() {
		for {
			select {
			case systemStatus := <-subscription:
				fmt.Println("\n=== TELEMETRÍA RECIBIDA ===")
				fmt.Printf("- Tiempo de misión: %s\n", systemStatus.FlightTelemetry.MissionTime)
				fmt.Printf("- Número de paquete: %d\n", systemStatus.FlightTelemetry.PacketCount)
				fmt.Printf("- Estado: %s\n", systemStatus.FlightTelemetry.Status)
				fmt.Printf("- Altitud: %.2f\n", systemStatus.FlightTelemetry.Altitude)
				fmt.Printf("- Inclinación: %.2f°\n", systemStatus.FlightTelemetry.Tilt)
				fmt.Printf("- GPS (Lat, Long, Alt): %.6f, %.6f, %.2f\n",
					systemStatus.FlightTelemetry.GPSLatitude,
					systemStatus.FlightTelemetry.GPSLongitude,
					systemStatus.FlightTelemetry.GPSAltitude)
				fmt.Printf("- Aceleración: %.2f\n", systemStatus.FlightTelemetry.Acceleration)
				fmt.Printf("- Temperatura: %.2f°C\n", systemStatus.FlightTelemetry.Temperature)
				fmt.Printf("- Voltaje de batería: %.2fV\n", systemStatus.FlightTelemetry.BatteryVoltage)
				fmt.Println("===========================")
			}
		}
	}()

	// Opción para seguir con el test hardcodeado si lo necesitas
	if len(os.Args) > 1 && os.Args[1] == "test" {
		fmt.Println("\nModo de prueba activado - Usando mensaje hardcodeado")
		testFlightTelemetry()
	}

	fmt.Println("\nEscuchando mensajes del XBee... (Ctrl+C para terminar)")

	// Esperar señal de interrupción
	sigChan := make(chan os.Signal, 1)
	signal.Notify(sigChan, syscall.SIGINT, syscall.SIGTERM)
	<-sigChan

	fmt.Println("\nCerrando conexión...")
}

// Función de prueba con mensaje hardcodeado (se mantiene por si la necesitas)
func testFlightTelemetry() {
	// Mensaje de telemetría hardcodeado (11 campos de 6 caracteres)
	testMessage := "000120000010NORMAL001275000450004258072500009800012500003500037500"

	fmt.Println("Usando mensaje de prueba:", testMessage)

	// Parsear el mensaje hardcodeado
	status, err := serial.ParseSystemStatus(testMessage)
	if err != nil {
		fmt.Printf("Error parseando telemetría: %v\n", err)
		return
	}

	// Imprimir los resultados del parseo
	fmt.Printf("\nResultado del parseo de telemetría:\n")
	fmt.Printf("- Tiempo de misión: %s\n", status.FlightTelemetry.MissionTime)
	fmt.Printf("- Número de paquete: %d\n", status.FlightTelemetry.PacketCount)
	fmt.Printf("- Estado: %s\n", status.FlightTelemetry.Status)
	fmt.Printf("- Altitud: %.2f\n", status.FlightTelemetry.Altitude)
	fmt.Printf("- Inclinación: %.2f°\n", status.FlightTelemetry.Tilt)
	fmt.Printf("- GPS (Lat, Long, Alt): %.6f, %.6f, %.2f\n",
		status.FlightTelemetry.GPSLatitude,
		status.FlightTelemetry.GPSLongitude,
		status.FlightTelemetry.GPSAltitude)
	fmt.Printf("- Aceleración: %.2f\n", status.FlightTelemetry.Acceleration)
	fmt.Printf("- Temperatura: %.2f°C\n", status.FlightTelemetry.Temperature)
	fmt.Printf("- Voltaje de batería: %.2fV\n", status.FlightTelemetry.BatteryVoltage)
}

/*
func main() {
	fmt.Println("=== INICIANDO PRUEBA DE TELEMETRÍA DE VUELO ===")

	// Imprimimos que vamos a comenzar la prueba
	fmt.Println("Ejecutando prueba con mensaje hardcodeado...")

	// Ejecutar la prueba con mensaje hardcodeado
	testFlightTelemetry()

	// Uso explícito del paquete model para evitar error de importación
	var _ = model.FlightTelemetrySystemStatus{}

	fmt.Println("=== PRUEBA FINALIZADA ===")
}

func testFlightTelemetry() {
	// Mensaje de telemetría hardcodeado (11 campos de 6 caracteres)
	// Formato: Tiempo,Paquete,Status,Alt,Tilt,LatGPS,LongGPS,AltGPS,Accel,Temp,Voltaje
	testMessage := "000120000010NORMAL001275000450004258072500009800012500003500037500"

	fmt.Println("Usando mensaje de prueba:", testMessage)

	// Parsear el mensaje hardcodeado
	status, err := serial.ParseSystemStatus(testMessage)
	if err != nil {
		fmt.Printf("Error parseando telemetría: %v\n", err)
		return
	}

	// Imprimir los resultados del parseo
	fmt.Printf("\nResultado del parseo de telemetría:\n")
	fmt.Printf("- Tiempo de misión: %s\n", status.FlightTelemetry.MissionTime)
	fmt.Printf("- Número de paquete: %d\n", status.FlightTelemetry.PacketCount)
	fmt.Printf("- Estado: %s\n", status.FlightTelemetry.Status)
	fmt.Printf("- Altitud: %.2f\n", status.FlightTelemetry.Altitude)
	fmt.Printf("- Inclinación: %.2f°\n", status.FlightTelemetry.Tilt)
	fmt.Printf("- GPS (Lat, Long, Alt): %.6f, %.6f, %.2f\n",
		status.FlightTelemetry.GPSLatitude,
		status.FlightTelemetry.GPSLongitude,
		status.FlightTelemetry.GPSAltitude)
	fmt.Printf("- Aceleración: %.2f\n", status.FlightTelemetry.Acceleration)
	fmt.Printf("- Temperatura: %.2f°C\n", status.FlightTelemetry.Temperature)
	fmt.Printf("- Voltaje de batería: %.2fV\n", status.FlightTelemetry.BatteryVoltage)
}
*/

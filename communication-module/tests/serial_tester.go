package main

import (
	"fmt"
	"kranz/communication-module/model"
	"kranz/communication-module/serial"
)

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

/*
Función de prueba para mensajes de launchpad comentada, ya que no se utiliza para la telemetría de vuelo
Si se necesita en el futuro, simplemente quitar los comentarios

func testLaunchpadMessage() {
	// Crea un mensaje simulado de la plataforma de lanzamiento
	testMessage := "STBY25.451.21.225.231true1true1000"

	// Intenta parsear el mensaje
	status, err := serial.ParseSystemStatus(testMessage)
	if err != nil {
		fmt.Printf("Error parseando mensaje de launchpad: %v\n", err)
		return
	}

	// Imprime el resultado
	fmt.Printf("Resultado del parseo de launchpad:\n")
	fmt.Printf("- Estado actual: %s\n", status.Launchpad.CurrentState)
	fmt.Printf("- Temperatura de ventilación: %.2f°C\n", status.OnBoard.TankDepressVentTempCelsius)
	fmt.Printf("- Presión de línea de carga: %.2f bar\n", status.Launchpad.LoadingLinePressureBar)
	fmt.Printf("- Presión de tierra: %.2f bar\n", status.Launchpad.GroundPressureBar)
	fmt.Printf("- Temperatura de tierra: %.2f°C\n", status.Launchpad.GroundTempCelsius)
	fmt.Printf("- Voltaje de batería OBEC: %.2fV\n", status.OnBoard.OBECBatteryVoltageVolt)
	fmt.Printf("- Válvula del motor abierta: %v\n", status.OnBoard.EngineValveOpen)
	fmt.Printf("- Válvula de carga abierta: %v\n", status.Launchpad.LoadingValveOpen)
	fmt.Printf("- Umbilical conectado: %v\n", status.Launchpad.UmbrilicalConnected)
	fmt.Printf("- Velocidad del viento: %d knt\n", status.WeatherData.WindSpeedKnt)
}
*/

package serial

import (
	"bufio"
	"errors"
	"fmt"
	"io"
	"kranz/communication-module/broadcast"
	"kranz/communication-module/model"
	"log"
	"os"
	"strconv"
	"sync"
	"time"

	"go.bug.st/serial"
)

var timeOfLastMessage = time.Now()
var latestSystemStatus = model.SystemStatus{}

func OpenSerialConnection(serialPort string) (serial.Port, sync.Mutex, error) {
	// Configuration for the serial port
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	mutex := sync.Mutex{}

	// Open the serial port (adjust "/dev/ttyUSB0" to your port)
	port, err := serial.Open(serialPort, mode)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
		return nil, mutex, errors.New(fmt.Sprintf("Failed to open serial port: %v", err))
	}
	return port, mutex, nil
}

func ListenForMessages(port serial.Port, portMutex *sync.Mutex, systemStatusBroadcast *broadcast.Broadcast[model.SystemStatus]) {
	// Buffer to store incoming data

	//Este caso es para MacOS
	//file, err := os.OpenFile("logs/"+time.Now().String()+"_telemetry_message_log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	//Este es para Windows +time.Now().String()+
	file, err := os.OpenFile("C:\\KranzMisssionControl\\kranz\\communication-module\\_telemetry_message_log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)

	if err != nil {
		fmt.Println("Error opening file:", err)
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("Telemetry message log for " + time.Now().String() + "\n")

	var buf []byte
	// Read from the port in a loop
	for {
		portMutex.Lock()
		// Read one byte
		b := make([]byte, 1)
		_, err := port.Read(b)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("port.Read: %v", err)
			}
		}
		portMutex.Unlock()

		// Append the byte to the buffer
		if b[0] != 0 {
			buf = append(buf, b[0])
		}

		// Check for the end character, e.g., newline ('\n')
		if b[0] == '\n' {
			fmt.Printf("Received message from XBEE: " + string(buf))
			timeOfLastMessage = time.Now()
			writer.WriteString(time.Now().String() + " ---> " + string(buf))
			parsed_status, err := ParseSystemStatus(string(buf[:len(buf)-1]))
			if err != nil {
				log.Println(err)
			}
			fmt.Printf("Parsed message to system status: %v\n", parsed_status)
			latestSystemStatus = parsed_status
			systemStatusBroadcast.SendBroadcast(parsed_status)
			buf = buf[:0]
			err = writer.Flush()
			if err != nil {
				fmt.Println("Error flushing data to file:", err)
				return
			}
		}

		/* Esto desconecta si pasaron más de 4 segundos sin mensajes

		if time.Since(timeOfLastMessage).Milliseconds() > 4000 {
			var systemStatusWithConnectionError = latestSystemStatus
			// Mantenemos esta línea por compatibilidad con el controlador
			systemStatusWithConnectionError.Launchpad = map[string]interface{}{
				"ConnectionStatus": "Last message received " + fmt.Sprintf("%f", time.Since(timeOfLastMessage).Seconds()) + " seconds ago",
			}
			systemStatusBroadcast.SendBroadcast(systemStatusWithConnectionError)
		}
		*/
	}
}

// Función SendCommand se mantiene comentada por si es necesaria en el futuro

func SendCommand(port serial.Port, portMutex *sync.Mutex, command model.Command) {
	portMutex.Lock()
	fmt.Println("Sending command: " + string(command.ToMessage()))
	port.Write(command.ToMessage())
	portMutex.Unlock()
}

func ParseSystemStatus(message string) (model.SystemStatus, error) {
	// Verificamos si estamos recibiendo un mensaje de telemetría (84 caracteres para 14 campos de 6 chars)
	if len(message) == 84 {
		return ParseFlightTelemetry(message)
	}

	// Si no es un mensaje de telemetría reconocible, devolvemos un error
	return model.SystemStatus{}, errors.New("message format not recognized or message is too short")
}

func ParseFlightTelemetry(message string) (model.SystemStatus, error) {
	// For flight telemetry, each field is exactly 6 characters
	const fieldSize = 6

	// Check if the message length is correct (14 fields * 6 chars)
	if len(message) != 14*fieldSize {
		return model.SystemStatus{}, errors.New(fmt.Sprintf("invalid telemetry message length: %d, expected: %d", len(message), 14*fieldSize))
	}

	// Helper function to parse float values from a 6-char field
	parseF32 := func(fieldValue string) (float32, error) {
		val, err := strconv.ParseFloat(fieldValue, 32)
		if err != nil {
			return 0, err
		}
		return float32(val), nil
	}

	// Extract all fields from the message
	fields := make([]string, 14)
	for i := 0; i < 14; i++ {
		start := i * fieldSize
		end := start + fieldSize
		fields[i] = message[start:end]
	}

	// Parse field values in the new order
	// 0: MISSION_TIME
	missionTime := fields[0]

	// 1: PACKET_COUNT
	packetCount, err := strconv.Atoi(fields[1])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse packet count")
	}

	// 2: STATUS
	status := fields[2]

	// 3: BATTERY_LEVEL
	batteryLevel, err := parseF32(fields[3])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse battery level")
	}

	// 4: IMU_Y_VEL
	imuYVel, err := parseF32(fields[4])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse IMU Y velocity")
	}

	// 5: IMU_ROLL
	imuRoll, err := parseF32(fields[5])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse IMU roll")
	}

	// 6: IMU_PITCH
	imuPitch, err := parseF32(fields[6])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse IMU pitch")
	}

	// 7: GNSS_TIME
	gnssTime := fields[7]

	// 8: GNSS_LATITUDE
	gnssLatitude, err := parseF32(fields[8])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse GNSS latitude")
	}

	// 9: GNSS_LONGITUDE
	gnssLongitude, err := parseF32(fields[9])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse GNSS longitude")
	}

	// 10: GNSS_ALTITUDE
	gnssAltitude, err := parseF32(fields[10])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse GNSS altitude")
	}

	// 11: BME_PRESSURE
	bmePressure, err := parseF32(fields[11])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse BME pressure")
	}

	// 12: BME_ALTITUDE
	bmeAltitude, err := parseF32(fields[12])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse BME altitude")
	}

	// 13: BME_TEMPERATURE
	bmeTemperature, err := parseF32(fields[13])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse BME temperature")
	}

	// Create a new system status with the telemetry data
	return model.SystemStatus{
		FlightTelemetry: model.FlightTelemetrySystemStatus{
			MissionTime:    missionTime,
			PacketCount:    packetCount,
			Status:         status,
			BatteryLevel:   batteryLevel,
			IMUYVel:        imuYVel,
			IMURoll:        imuRoll,
			IMUPitch:       imuPitch,
			GNSSTime:       gnssTime,
			GNSSLatitude:   gnssLatitude,
			GNSSLongitude:  gnssLongitude,
			GNSSAltitude:   gnssAltitude,
			BMEPressure:    bmePressure,
			BMEAltitude:    bmeAltitude,
			BMETemperature: bmeTemperature,
		},
		// Dejamos vacíos los demás campos que ya no se usan
		OnBoard:     map[string]interface{}{"ConnectionStatus": "OK"},
		Launchpad:   map[string]interface{}{"ConnectionStatus": "OK"},
		WeatherData: map[string]interface{}{},
	}, nil
}

/*
Funciones específicas de LaunchPad comentadas, ya que no se utilizan para la telemetría de vuelo
Si se necesitan en el futuro, simplemente quitar los comentarios

// Original parsing function renamed for backward compatibility
func ParseLaunchpadStatus(message string) (model.SystemStatus, error) {
	if len(message) < 35 {
		return model.SystemStatus{}, errors.New("message is too short")
	}

	parseF32 := func(s string) (float32, error) {
		val, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return 0, err
		}
		return float32(val), nil
	}

	parseBool := func(s string) (bool, error) {
		val, err := strconv.ParseBool(s)
		if err != nil {
			return false, err
		}
		return val, nil
	}

	lcState, err := ParseLCState(message[0:4])
	if err != nil {
		return model.SystemStatus{}, err
	}

	tankDepressVentTempCelsius, err := parseF32(message[4:9])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse tank depress vent temperature from SystemMessage")
	}

	loadingLinePressureBar, err := parseF32(message[9:13])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse loading line pressure from SystemMessage")
	}

	groundPressureBar, err := parseF32(message[13:17])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse ground pressure from SystemMessage")
	}

	groundTempCelsius, err := parseF32(message[17:22])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse ground temperature from SystemMessage")
	}

	obecBatteryVoltageVolt, err := parseF32(message[22:27])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse OBEC battery voltage from SystemMessage")
	}

	obecConnectionOK, err := parseBool(message[27:28])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse OBEC Connection OK from SystemMessage")
	}
	engineValveOpen, err := parseBool(message[28:29])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse engine valve open from SystemMessage")
	}
	loadingValveOpen, err := parseBool(message[29:30])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse loading valve open from SystemMessage")
	}
	umbrilicalConnected, err := parseBool(message[30:31])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse umbrilical connected from SystemMessage")
	}
	umbrilicalFinishedDisconnect, err := parseBool(message[31:32])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse umbrilical finished disconnect from SystemMessage")
	}
	windSpeedKnt, err := strconv.ParseUint(message[32:35], 10, 8)
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse wind speed from SystemMessage")
	}

	obecConnectionOKString := ""
	if obecConnectionOK {
		obecConnectionOKString = "OK"
	} else {
		obecConnectionOKString = "ERROR"
	}

	return model.SystemStatus{
		// Initialize an empty FlightTelemetry structure
		FlightTelemetry: model.FlightTelemetrySystemStatus{},
		OnBoard: model.OnBoardSystemStatus{
			ConnectionStatus:           obecConnectionOKString,
			TankDepressVentTempCelsius: tankDepressVentTempCelsius,
			EngineValveOpen:            engineValveOpen,
			OBECBatteryVoltageVolt:     obecBatteryVoltageVolt,
			FlightComputersStatus:      model.FlightComputersStatus{AdaOK: true, AltiumOK: true},
		},
		Launchpad: model.LaunchpadSystemStatus{
			CurrentState:                 lcState,
			ConnectionStatus:             "Ok",
			LoadingLinePressureBar:       loadingLinePressureBar,
			GroundPressureBar:            groundPressureBar,
			GroundTempCelsius:            groundTempCelsius,
			LoadingValveOpen:             loadingValveOpen,
			UmbrilicalConnected:          umbrilicalConnected,
			UmbrilicalFinishedDisconnect: umbrilicalFinishedDisconnect,
		},
		WeatherData: model.WeatherData{
			WindSpeedKnt: uint8(windSpeedKnt),
		},
	}, nil
}

func ParseLCState(str string) (model.LCState, error) {
	switch str {
	case "STBY":
		return model.STANDBY, nil
	case "STPW":
		return model.STANDBY_PRESSURE_WARNING, nil
	case "LDNG":
		return model.LOADING, nil
	case "PRFC":
		return model.PRE_FLIGHT_CHECK, nil
	case "PRLW":
		return model.PRE_LAUNCH_WIND_CHECK, nil
	case "PRLU":
		return model.PRE_LAUNCH_UMBRILICAL_DISCONNECT, nil
	case "IGVO":
		return model.IGNITION_OPEN_VALVE, nil
	case "IGON":
		return model.IGNITION_IGNITERS_ON, nil
	case "IGOF":
		return model.IGNITION_IGNITERS_OFF, nil
	case "ABRT":
		return model.ABORT, nil
	default:
		return model.STANDBY, errors.New("LCState string does not match any state: " + str)
	}
}
*/

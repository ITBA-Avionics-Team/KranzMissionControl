package serial

import (
	"errors"
	"fmt"
	"io"
	"kranz/communication-module/broadcast"
	"kranz/communication-module/model"
	"log"
	"strconv"
	"sync"

	"go.bug.st/serial"
)

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
	// buf := make([]byte, 128)

	// Read from the port in a loop
	for {
		portMutex.Lock()
		var buf []byte
		for {
			// Read one byte
			b := make([]byte, 1)
			_, err := port.Read(b)
			if err != nil {
				if err != io.EOF {
					log.Fatalf("port.Read: %v", err)
				}
				portMutex.Unlock()
				break
			}

			// Append the byte to the buffer
			buf = append(buf, b[0])
			// fmt.Printf(string(buf));

			// Check for the end character, e.g., newline ('\n')
			if b[0] == '\n' {
				fmt.Printf("Received message from XBEE: " + string(buf))
				parsed_status, err := ParseSystemStatus(string(buf[:len(buf)-1]))
				if err != nil {
					log.Println(err)
				}
				fmt.Printf("Parsed message to system status: %v", parsed_status)
				systemStatusBroadcast.SendBroadcast(parsed_status)
				buf = buf[:0]
				break // End character found, exit the loop
			}
		}
		portMutex.Unlock()
	}
}

func SendCommand(port serial.Port, portMutex *sync.Mutex, command model.Command) {
	portMutex.Lock()
	fmt.Println("Sending command: " + string(command.ToMessage()))
	port.Write(command.ToMessage())
	portMutex.Unlock()
}

func ParseSystemStatus(message string) (model.SystemStatus, error) {
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

	loadLinePressureBar, err := parseF32(message[9:13])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse load line pressure from SystemMessage")
	}

	groundTempCelsius, err := parseF32(message[13:18])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse ground temperature from SystemMessage")
	}

	obecBatteryVoltageVolt, err := parseF32(message[18:23])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse OBEC battery voltage from SystemMessage")
	}

	obecConnectionOK, err := parseBool(message[23:24])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse OBEC Connection OK from SystemMessage")
	}
	engineValveOpen, err := parseBool(message[24:25])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse engine valve open from SystemMessage")
	}
	loadingValveOpen, err := parseBool(message[25:26])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse loading valve open from SystemMessage")
	}
	loadingDepressVentValveOpen, err := parseBool(message[26:27])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse loading depress vent valve open from SystemMessage")
	}
	umbrilicalConnected, err := parseBool(message[27:28])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse umbrilical connected from SystemMessage")
	}
	umbrilicalFinishedDisconnect, err := parseBool(message[28:28])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse umbrilical finished disconnect from SystemMessage")
	}
	igniterContinuityOK, err := parseBool(message[28:29])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse Igniter Continuity OK from SystemMessage")
	}
	externalVentAsDefault, err := parseBool(message[29:30])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse external vent as default from SystemMessage")
	}
	windSpeedKnt, err := strconv.ParseUint(message[30:33], 10, 8)
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
		OnBoard: model.OnBoardSystemStatus{
			ConnectionStatus:           obecConnectionOKString,
			TankDepressVentTempCelsius: tankDepressVentTempCelsius,
			EngineValveOpen:            engineValveOpen,
			OBECBatteryVoltageVolt:     obecBatteryVoltageVolt,
			FlightComputersStatus:      model.FlightComputersStatus{AdaOK: true, AltiumOK: true},
		},
		Launchpad: model.LaunchpadSystemStatus{
			CurrentState:                 lcState,
			ConnectionStatus:             "ok",
			LoadLinePressureBar:          loadLinePressureBar,
			GroundTempCelsius:            groundTempCelsius,
			LoadingValveOpen:             loadingValveOpen,
			LoadingDepressVentValveOpen:  loadingDepressVentValveOpen,
			UmbrilicalConnected:          umbrilicalConnected,
			UmbrilicalFinishedDisconnect: umbrilicalFinishedDisconnect,
			IgniterContinuityOK:          igniterContinuityOK,
			ExternalVentAsDefault:        externalVentAsDefault,
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

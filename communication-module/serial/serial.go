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


func OpenSerialConnection(serialPort string) (serial.Port, sync.Mutex, error){
	// Configuration for the serial port
	mode := &serial.Mode{
		BaudRate: 115200,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	mutex := sync.Mutex{};

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
				fmt.Printf(string(buf));
				parsed_status, err := ParseSystemStatus(string(buf[:len(buf)-1]));
				if err != nil {
					log.Println(err);
					fmt.Printf(string(buf));
				}
				systemStatusBroadcast.SendBroadcast(parsed_status)
				buf = buf[:0]
				break // End character found, exit the loop
			}
		}
		portMutex.Unlock();
	}
}

func SendCommand(port serial.Port, portMutex *sync.Mutex, command model.Command) {
	portMutex.Lock()
	fmt.Println("Sending command: " + string(command.ToMessage()))
	port.Write(command.ToMessage())
	portMutex.Unlock()
}

func ParseSystemStatus(message string) (model.SystemStatus, error) {
	if len(message) < 32 {
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

	tankPressurePSI, err := parseF32(message[4:8])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse tank pressure from SystemMessage")
	}

	tankTempCelsius, err := parseF32(message[8:12])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse tank temperature from SystemMessage")
	}

	tankDepressVentTempCelsius, err := parseF32(message[12:16])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse tank depress vent temperature from SystemMessage")
	}

	loadLinePressurePsi, err := parseF32(message[16:20])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse load line pressure from SystemMessage")
	}

	obecBatteryVoltage, err := parseF32(message[20:24])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse OBEC battery voltage from SystemMessage")
	}

	lcBatteryVoltage, err := parseF32(message[24:28])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse LC battery voltage from SystemMessage")
	}

	obecConnectionOK, err := parseBool(message[28:29])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse OBEC Connection OK from SystemMessage")
	}
	tankDepressVentValveOpen, err := parseBool(message[29:30])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse tank depress vent valve open from SystemMessage")
	}
	engineValveOpen, err := parseBool(message[30:31])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse engine valve open from SystemMessage")
	}
	loadingValveOpen, err := parseBool(message[31:32])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse loading valve open from SystemMessage")
	}
	loadingDepressVentValveOpen, err := parseBool(message[32:33])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse loading depress vent valve open from SystemMessage")
	}
	umbrilicalConnected, err := parseBool(message[33:34])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse umbrilical connected from SystemMessage")
	}
	igniterContinuityOK, err := parseBool(message[34:35])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse Igniter Continuity OK from SystemMessage")
	}
	externalVentAsDefault, err := parseBool(message[35:36])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse external vent as default from SystemMessage")
	}
	windSpeedKnt, err := strconv.ParseUint(message[36:39], 10, 8)
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse wind speed from SystemMessage")
	}

	obecConnectionOKString := ""
	if (obecConnectionOK){
		obecConnectionOKString = "OK"
	} else {
		obecConnectionOKString = "ERROR"
	}

	return model.SystemStatus{
		OnBoard: model.OnBoardSystemStatus{
			ConnectionStatus:           obecConnectionOKString,
			TankPressurePSI:            tankPressurePSI,
			TankTempCelsius:            tankTempCelsius,
			TankDepressVentTempCelsius: tankDepressVentTempCelsius,
			TankDepressVentValveOpen:   tankDepressVentValveOpen,
			EngineValveOpen:            engineValveOpen,
			OBECBatteryVoltageVolt:     obecBatteryVoltage,
			FlightComputersStatus:      model.FlightComputersStatus{AdaOK: true, AltiumOK: true},
		},
		Launchpad: model.LaunchpadSystemStatus{
			CurrentState:                lcState,
			ConnectionStatus:            "ok",
			LoadLinePressurePsi:         loadLinePressurePsi,
			LoadingValveOpen:            loadingValveOpen,
			LoadingDepressVentValveOpen: loadingDepressVentValveOpen,
			UmbrilicalConnected:         umbrilicalConnected,
			IgniterContinuityOK:         igniterContinuityOK,
			ExternalVentAsDefault:       externalVentAsDefault,
			LCBatteryVoltageVolt:        lcBatteryVoltage,
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
	case "STPE":
		return model.STANDBY_PRESSURE_WARNING_EXTERNAL_VENT, nil
	case "LDNG":
		return model.LOADING, nil
	case "PRFC":
		return model.PRE_FLIGHT_CHECK, nil
	case "PRLW":
		return model.PRE_LAUNCH_WIND_CHECK, nil
	case "PRLU":
		return model.PRE_LAUNCH_UMBRILICAL_DISCONNECT, nil
	case "IGON":
		return model.IGNITION_IGNITERS_ON, nil
	case "IGVO":
		return model.IGNITION_OPEN_VALVE, nil
	case "IGOF":
		return model.IGNITION_IGNITERS_OFF, nil
	case "ABRT":
		return model.ABORT, nil
	default:
		return model.STANDBY, errors.New("LCState string does not match any state: " + str)
	}
}

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

	file, err := os.OpenFile("logs/"+time.Now().String()+"_LC_message_log.txt", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	writer := bufio.NewWriter(file)
	writer.WriteString("LC message log for " + time.Now().String() + "\n")

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

		if time.Since(timeOfLastMessage).Milliseconds() > 4000 {
			var systemStatusWithLCConnectionError = latestSystemStatus
			systemStatusWithLCConnectionError.Launchpad.ConnectionStatus = "Last message received " + fmt.Sprintf("%f", time.Since(timeOfLastMessage).Seconds()) + " seconds ago"
			systemStatusBroadcast.SendBroadcast(systemStatusWithLCConnectionError)
		}

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

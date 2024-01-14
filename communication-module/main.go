package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"log"
	"strconv"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/websocket"
	"go.bug.st/serial"
)

var upgrader = websocket.Upgrader{
	ReadBufferSize:  1024,
	WriteBufferSize: 1024,
}

func main() {
	systemStatusChannel := make(chan SystemStatus)
	go listen_for_messages(systemStatusChannel)
	router := gin.Default()
	router.GET("/system_status", func(c *gin.Context) {
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			return
		}
		defer conn.Close()
		for {
			select {
			case systemStatusMessage := <- systemStatusChannel:
				fmt.Printf("[Controller] Sending message: %+v\n", systemStatusMessage)
				jsonData, err := json.Marshal(systemStatusMessage)
				if err != nil {
					log.Fatalf("Error marshaling to JSON: %v", err)
				}
				conn.WriteMessage(websocket.TextMessage, jsonData)
				break
			}
		}
	})
	router.Run(":8080")
}

func listen_for_messages(systemStatusChannel chan SystemStatus) {
	// Configuration for the serial port
	mode := &serial.Mode{
		BaudRate: 9600,
		Parity:   serial.NoParity,
		DataBits: 8,
		StopBits: serial.OneStopBit,
	}

	// Open the serial port (adjust "/dev/ttyUSB0" to your port)
	port, err := serial.Open("/dev/ttys008", mode)
	if err != nil {
		log.Fatalf("serial.Open: %v", err)
	}
	defer port.Close()

	// Buffer to store incoming data
	buf := make([]byte, 128)

	// Read from the port in a loop
	for {
		n, err := port.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("port.Read: %v", err)
			}
			break
		}
		// Process the incoming data
		fmt.Printf("%s", string(buf[:n]))
		parsed_status, err := ParseSystemStatus(string(buf[:n]));
		fmt.Printf("%v", parsed_status)
		if err != nil {
			fmt.Printf("Failed to parse system status :(");
		}
		systemStatusChannel <- parsed_status
	}
}

type FlightComputersStatus struct {
	AltiumOK bool `json:"altium_ok"`
	AdaOK    bool `json:"ada_ok"`
}

type OnBoardSystemStatus struct {
	ConnectionStatus           string                `json:"connection_status"`
	TankPressurePSI            float32               `json:"tank_pressure_psi"`
	TankTempCelsius            float32               `json:"tank_temp_celsius"`
	TankDepressVentTempCelsius float32               `json:"tank_depress_vent_temp_celsius"`
	TankDepressVentValveOpen   bool                  `json:"tank_depress_vent_valve_open"`
	EngineValveOpen            bool                  `json:"engine_valve_open"`
	OBECBatteryVoltageVolt     float32               `json:"obec_battery_voltage_volt"`
	FlightComputersStatus      FlightComputersStatus `json:"flight_computers_status"`
}

type LCState string

const (
	STANDBY                                LCState = "STANDBY"
	STANDBY_PRESSURE_WARNING               LCState = "STANDBY_PRESSURE_WARNING"
	STANDBY_PRESSURE_WARNING_EXTERNAL_VENT LCState = "STANDBY_PRESSURE_WARNING_EXTERNAL_VENT"
	LOADING                                LCState = "LOADING"
	PRE_FLIGHT_CHECK                       LCState = "PRE_FLIGHT_CHECK"
	PRE_LAUNCH_WIND_CHECK                  LCState = "PRE_LAUNCH_WIND_CHECK"
	PRE_LAUNCH_UMBRILICAL_DISCONNECT       LCState = "PRE_LAUNCH_UMBRILICAL_DISCONNECT"
	IGNITION_IGNITERS_ON                   LCState = "IGNITION_IGNITERS_ON"
	IGNITION_OPEN_VALVE                    LCState = "IGNITION_OPEN_VALVE"
	IGNITION_IGNITERS_OFF                  LCState = "IGNITION_IGNITERS_OFF"
	ABORT                                  LCState = "ABORT"
)

type LaunchpadSystemStatus struct {
	CurrentState                LCState `json:"current_state`
	ConnectionStatus            string  `json:"connection_status"`
	LoadLinePressurePsi         float32 `json:"load_line_pressure_psi"`
	LoadingValveOpen            bool    `json:"loading_valve_open"`
	LoadingDepressVentValveOpen bool    `json:"loading_depress_vent_valve_open"`
	UmbrilicalConnected         bool    `json:"umbrilical_connected"`
	IgniterContinuityOK         bool    `json:"igniter_continuity_ok"`
	ExternalVentAsDefault       bool    `json:"external_vent_as_default"`
	LCBatteryVoltageVolt        float32 `json:"lc_battery_voltage_volt"`
}

type WeatherData struct {
	WindSpeedKnt uint8 `json:"wind_speed_knt"`
}

type SystemStatus struct {
	OnBoard     OnBoardSystemStatus   `json:"on_board"`
	Launchpad   LaunchpadSystemStatus `json:"launchpad"`
	WeatherData WeatherData           `json:"weather_data"`
}

func ParseSystemStatus(message string) (SystemStatus, error) {
	if len(message) < 32 {
		return SystemStatus{}, errors.New("message is too short")
	}

	sensorDataByte := message[28]

	// obecConnectionOK := sensorDataByte&0b01 == 1
	tankDepressVentValveOpen := sensorDataByte&0b10 == 1
	engineValveOpen := sensorDataByte&0b100 == 1
	loadingValveOpen := sensorDataByte&0b1000 == 1
	loadingDepressVentValveOpen := sensorDataByte&0b10000 == 1
	umbrilicalConnected := sensorDataByte&0b100000 == 1
	igniterContinuityOK := sensorDataByte&0b1000000 == 1
	externalVentAsDefault := sensorDataByte&0b10000000 == 1

	parseF32 := func(s string) (float32, error) {
		val, err := strconv.ParseFloat(s, 32)
		if err != nil {
			return 0, err
		}
		return float32(val), nil
	}

	lcState, err := ParseLCState(message[0:4])
	if err != nil {
		return SystemStatus{}, err
	}

	tankPressurePSI, err := parseF32(message[4:8])
	if err != nil {
		return SystemStatus{}, errors.New("failed to parse tank pressure from SystemMessage")
	}

	tankTempCelsius, err := parseF32(message[8:12])
	if err != nil {
		return SystemStatus{}, errors.New("failed to parse tank temperature from SystemMessage")
	}

	tankDepressVentTempCelsius, err := parseF32(message[12:16])
	if err != nil {
		return SystemStatus{}, errors.New("failed to parse tank depress vent temperature from SystemMessage")
	}

	obecBatteryVoltage, err := parseF32(message[20:24])
	if err != nil {
		return SystemStatus{}, errors.New("failed to parse OBEC battery voltage from SystemMessage")
	}

	loadLinePressurePsi, err := parseF32(message[16:20])
	if err != nil {
		return SystemStatus{}, errors.New("failed to parse load line pressure from SystemMessage")
	}

	lcBatteryVoltage, err := parseF32(message[24:28])
	if err != nil {
		return SystemStatus{}, errors.New("failed to parse LC battery voltage from SystemMessage")
	}

	windSpeedKnt, err := strconv.ParseUint(message[29:32], 10, 8)
	if err != nil {
		return SystemStatus{}, errors.New("failed to parse wind speed from SystemMessage")
	}

	return SystemStatus{
		OnBoard: OnBoardSystemStatus{
			ConnectionStatus:           "ok",
			TankPressurePSI:            tankPressurePSI,
			TankTempCelsius:            tankTempCelsius,
			TankDepressVentTempCelsius: tankDepressVentTempCelsius,
			TankDepressVentValveOpen:   tankDepressVentValveOpen,
			EngineValveOpen:            engineValveOpen,
			OBECBatteryVoltageVolt:     obecBatteryVoltage,
			FlightComputersStatus:      FlightComputersStatus{AdaOK: true, AltiumOK: true},
		},
		Launchpad: LaunchpadSystemStatus{
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
		WeatherData: WeatherData{
			WindSpeedKnt: uint8(windSpeedKnt),
		},
	}, nil
}

func ParseLCState(str string) (LCState, error) {
	switch str {
	case "STBY":
		return STANDBY, nil
	case "STPW":
		return STANDBY_PRESSURE_WARNING, nil
	case "STPE":
		return STANDBY_PRESSURE_WARNING_EXTERNAL_VENT, nil
	case "LDNG":
		return LOADING, nil
	case "PRFC":
		return PRE_FLIGHT_CHECK, nil
	case "PRLW":
		return PRE_LAUNCH_WIND_CHECK, nil
	case "PRLU":
		return PRE_LAUNCH_UMBRILICAL_DISCONNECT, nil
	case "IGON":
		return IGNITION_IGNITERS_ON, nil
	case "IGVO":
		return IGNITION_OPEN_VALVE, nil
	case "IGOF":
		return IGNITION_IGNITERS_OFF, nil
	case "ABRT":
		return ABORT, nil
	default:
		return STANDBY, errors.New("LCState string does not match any state: " + str)
	}
}

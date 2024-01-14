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
		BaudRate: 9600,
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
	buf := make([]byte, 128)

	// Read from the port in a loop
	for {
		portMutex.Lock()
		n, err := port.Read(buf)
		if err != nil {
			if err != io.EOF {
				log.Fatalf("port.Read: %v", err)
			}
			portMutex.Unlock()
			break
		}
		// Process the incoming data
		fmt.Printf("%s", string(buf[:n]))
		parsed_status, err := ParseSystemStatus(string(buf[:n]));
		fmt.Printf("%v", parsed_status)
		if err != nil {
			fmt.Printf("Failed to parse system status :(");
		}
		systemStatusBroadcast.SendBroadcast(parsed_status)
		portMutex.Unlock()
	}
}

func SendCommand(port serial.Port, portMutex *sync.Mutex, command model.Command) {
	portMutex.Lock()
	port.Write(command.ToMessage())
	portMutex.Unlock()
}

func ParseSystemStatus(message string) (model.SystemStatus, error) {
	if len(message) < 32 {
		return model.SystemStatus{}, errors.New("message is too short")
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

	obecBatteryVoltage, err := parseF32(message[20:24])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse OBEC battery voltage from SystemMessage")
	}

	loadLinePressurePsi, err := parseF32(message[16:20])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse load line pressure from SystemMessage")
	}

	lcBatteryVoltage, err := parseF32(message[24:28])
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse LC battery voltage from SystemMessage")
	}

	windSpeedKnt, err := strconv.ParseUint(message[29:32], 10, 8)
	if err != nil {
		return model.SystemStatus{}, errors.New("failed to parse wind speed from SystemMessage")
	}

	return model.SystemStatus{
		OnBoard: model.OnBoardSystemStatus{
			ConnectionStatus:           "ok",
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




// String get_4_byte_string_from_valve(Valve valve)
// {
//   switch (valve)
//   {
//   case TANK_DEPRESS_VENT_VALVE:
//     return "TDVV";
//   case ENGINE_VALVE:
//     return "ENGV";
//   case LOADING_VALVE:
//     return "LDGV";
//   case LOADING_LINE_DEPRESS_VENT_VALVE:
//     return "LDVV";
//   }
// }


// static String to_message(Command command)
//   {
//     switch (command.type)
//     {
//     case ValveCommand:
//       switch (command.valve)
//       {
//       case TANK_DEPRESS_VENT_VALVE:
//         return String("VCTDVV") + String((char)command.uint_value, 1) + String("|");
//         break;
//       case ENGINE_VALVE:
//         return String("VCENGV") + String((char)command.uint_value, 1) + String("|");
//         break;
//       case LOADING_VALVE:
//         return String("VCLDGV") + String((char)command.uint_value, 1) + String("|");
//         break;
//       case LOADING_LINE_DEPRESS_VENT_VALVE:
//         return String("VCLDVV") + String((char)command.uint_value, 1) + String("|");
//         break;
//       }
//       break;
//     case SwitchStateCommand:
//       return String("SS") + get_4_byte_string_from_state(command.state) + String("|");
//       break;
//     case SetExternalVentAsDefaultCommand:
//       return String("EV") + command.bool_value ? String("1") : String("0") + String("|");
//       break;
//     case EMPTY:
//       break;
//     }
//   }

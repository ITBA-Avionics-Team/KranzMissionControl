package model

import "fmt"

/*
FLIGHT TELEMETRY MODELS
*/
type FlightTelemetrySystemStatus struct {
	MissionTime     string  `json:"mission_time"`
	PacketCount     int     `json:"packet_count"`
	Status          string  `json:"status"`
	Altitude        float32 `json:"altitude"`
	Tilt            float32 `json:"tilt"`
	GPSLatitude     float32 `json:"gps_latitude"`
	GPSLongitude    float32 `json:"gps_longitude"`
	GPSAltitude     float32 `json:"gps_altitude"`
	Acceleration    float32 `json:"acceleration"`
	Temperature     float32 `json:"temperature"`
	BatteryVoltage  float32 `json:"battery_voltage"`
}

// Estructura simplificada del sistema que solo contiene telemetría
type SystemStatus struct {
	FlightTelemetry FlightTelemetrySystemStatus `json:"flight_telemetry"`
	// Mantenemos estos campos vacíos por compatibilidad con el controller existente
	// pero no serán utilizados para la telemetría de vuelo
	OnBoard     interface{} `json:"on_board"`
	Launchpad   interface{} `json:"launchpad"`
	WeatherData interface{} `json:"weather_data"`
}

/*
Las siguientes secciones están comentadas porque son específicas del LaunchPad
y ya no se utilizan para la telemetría de vuelo.
Si necesitas volver a habilitarlas, simplemente quita los comentarios.
*/


// SYSTEM STATUS MODELS
type FlightComputersStatus struct {
	AltiumOK bool `json:"altium_ok"`
	AdaOK    bool `json:"ada_ok"`
}

type OnBoardSystemStatus struct {
	ConnectionStatus           string                `json:"connection_status"`
	TankDepressVentTempCelsius float32               `json:"tank_depress_vent_temp_celsius"`
	EngineValveOpen            bool                  `json:"engine_valve_open"`
	OBECBatteryVoltageVolt     float32               `json:"obec_battery_voltage_volt"`
	FlightComputersStatus      FlightComputersStatus `json:"flight_computers_status"`
}

type LCState string

const (
	STANDBY                          LCState = "STANDBY"
	STANDBY_PRESSURE_WARNING         LCState = "STANDBY_PRESSURE_WARNING"
	LOADING                          LCState = "LOADING"
	PRE_FLIGHT_CHECK                 LCState = "PRE_FLIGHT_CHECK"
	PRE_LAUNCH_WIND_CHECK            LCState = "PRE_LAUNCH_WIND_CHECK"
	PRE_LAUNCH_UMBRILICAL_DISCONNECT LCState = "PRE_LAUNCH_UMBRILICAL_DISCONNECT"
	IGNITION_OPEN_VALVE              LCState = "IGNITION_OPEN_VALVE"
	IGNITION_IGNITERS_ON             LCState = "IGNITION_IGNITERS_ON"
	IGNITION_IGNITERS_OFF            LCState = "IGNITION_IGNITERS_OFF"
	ABORT                            LCState = "ABORT"
)

func Get4ByteStringFromState(state LCState) string {
	switch state {
	case STANDBY:
		return "STBY"
	case STANDBY_PRESSURE_WARNING:
		return "STPW"
	case LOADING:
		return "LDNG"
	case PRE_FLIGHT_CHECK:
		return "PRFC"
	case PRE_LAUNCH_WIND_CHECK:
		return "PRLW"
	case PRE_LAUNCH_UMBRILICAL_DISCONNECT:
		return "PRLU"
	case IGNITION_OPEN_VALVE:
		return "IGVO"
	case IGNITION_IGNITERS_ON:
		return "IGON"
	case IGNITION_IGNITERS_OFF:
		return "IGOF"
	case ABORT:
		return "ABRT"
	}
	return ""
}

type LaunchpadSystemStatus struct {
	ConnectionStatus             string  `json:"connection_status"`
	CurrentState                 LCState `json:"current_state"`
	LoadingLinePressureBar       float32 `json:"loading_line_pressure_bar"`
	GroundPressureBar            float32 `json:"ground_pressure_bar"`
	GroundTempCelsius            float32 `json:"ground_temp_celsius"`
	LoadingValveOpen             bool    `json:"loading_valve_open"`
	LoadingDepressVentValveOpen  bool    `json:"loading_depress_vent_valve_open"`
	UmbrilicalConnected          bool    `json:"umbrilical_connected"`
	UmbrilicalFinishedDisconnect bool    `json:"umbrilical_finished_disconnect"`
	IgniterContinuityOK          bool    `json:"igniter_continuity_ok"`
	ExternalVentAsDefault        bool    `json:"external_vent_as_default"`
}

type WeatherData struct {
	WindSpeedKnt uint8 `json:"wind_speed_knt"`
}

// COMMAND MODELS
type CommandType string

const (
	VALVE_COMMAND                        CommandType = "VALVE_COMMAND"
	SWITCH_STATE_COMMAND                 CommandType = "SWITCH_STATE_COMMAND"
	SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND CommandType = "SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND"
	EMPTY_COMMAND                        CommandType = "EMPTY_COMMAND"
	RAW_COMMAND                          CommandType = "RAW_COMMAND"
)

type Valve string

const (
	ENGINE_VALVE  Valve = "ENGINE_VALVE"
	LOADING_VALVE Valve = "LOADING_VALVE"
)

type Command struct {
	CommandType CommandType `json:"command_type" binding:"required"`
	UintValue   uint        `json:"uint_value"`
	Valve       Valve       `json:"valve"`
	State       LCState     `json:"state"`
	BoolValue   bool        `json:"bool_value"`
	StringValue string      `json:"string_value"`
}

func (command *Command) ToMessage() []byte {
	switch command.CommandType {
	case VALVE_COMMAND:
		switch command.Valve {
		case ENGINE_VALVE:
			return []byte(fmt.Sprintf("VCENGV%c|", command.UintValue))
		case LOADING_VALVE:
			return []byte(fmt.Sprintf("VCLDGV%c|", command.UintValue))
		}
	case SWITCH_STATE_COMMAND:
		return []byte("SS" + Get4ByteStringFromState(command.State) + "|")
	case SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND:
		if command.BoolValue {
			return []byte("EV1|")
		}
		return []byte("EV0|")
	case EMPTY_COMMAND:
		// Do nothing for EMPTY
		return []byte{}
	case RAW_COMMAND:
		return []byte(command.StringValue)
	}
	return []byte{}
}


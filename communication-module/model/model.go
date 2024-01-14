package model

/*
SYSTEM STATUS MODELS
*/
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

/*
COMMAND MODELS
*/
type CommandType string

const (
	VALVE_COMMAND                        CommandType = "VALVE_COMMAND"
	SWITCH_STATE_COMMAND                 CommandType = "SWITCH_STATE_COMMAND"
	SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND CommandType = "SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND"
	EMPTY_COMMAND                        CommandType = "EMPTY_COMMAND"
)

type Valve string

const (
	TANK_DEPRESS_VENT_VALVE         Valve = "TANK_DEPRESS_VENT_VALVE"
	ENGINE_VALVE                    Valve = "ENGINE_VALVE"
	LOADING_VALVE                   Valve = "LOADING_VALVE"
	LOADING_LINE_DEPRESS_VENT_VALVE Valve = "LOADING_LINE_DEPRESS_VENT_VALVE"
)

type Command struct {
	command_type CommandType
	uint_value   uint
	valve        Valve
	state        LCState
	bool_value   bool
}

func (command *Command) ToMessage() []byte {
	return []byte("SSABRT") // TODO: Implement
}

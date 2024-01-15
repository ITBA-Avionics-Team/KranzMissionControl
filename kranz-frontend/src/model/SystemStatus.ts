interface FlightComputersStatus {
  altium_ok: boolean;
  ada_ok: boolean;
}

interface OnBoardSystemStatus {
  connection_status: string;
  tank_pressure_psi: number;
  tank_temp_celsius: number;
  tank_depress_vent_temp_celsius: number;
  tank_depress_vent_valve_open: boolean;
  engine_valve_open: boolean;
  obec_battery_voltage_volt: number;
  flight_computers_status: FlightComputersStatus;
}

type LCState = 'STANDBY' | 'STANDBY_PRESSURE_WARNING' | 'STANDBY_PRESSURE_WARNING_EXTERNAL_VENT' | 'LOADING' | 'PRE_FLIGHT_CHECK' | 'PRE_LAUNCH_WIND_CHECK' | 'PRE_LAUNCH_UMBRILICAL_DISCONNECT' | 'IGNITION_IGNITERS_ON' | 'IGNITION_OPEN_VALVE' | 'IGNITION_IGNITERS_OFF' | 'ABORT';

interface LaunchpadSystemStatus {
  current_state: LCState;
  connection_status: string;
  load_line_pressure_psi: number;
  loading_valve_open: boolean;
  loading_depress_vent_valve_open: boolean;
  umbrilical_connected: boolean;
  igniter_continuity_ok: boolean;
  external_vent_as_default: boolean;
  lc_battery_voltage_volt: number;
}

interface WeatherData {
  wind_speed_knt: number;
}

export interface SystemStatus {
  on_board: OnBoardSystemStatus;
  launchpad: LaunchpadSystemStatus;
  weather_data: WeatherData;
}

export const DefaultSystemStatus: SystemStatus = {
  on_board: {
    connection_status: "",
    tank_pressure_psi: 0,
    tank_temp_celsius: 0,
    tank_depress_vent_temp_celsius: 0,
    tank_depress_vent_valve_open: false,
    engine_valve_open: false,
    obec_battery_voltage_volt: 0,
    flight_computers_status: {
      altium_ok: false,
      ada_ok: false,
    },
  },
  launchpad: {
    current_state: 'STANDBY',
    connection_status: "",
    load_line_pressure_psi: 0,
    loading_valve_open: false,
    loading_depress_vent_valve_open: false,
    umbrilical_connected: false,
    igniter_continuity_ok: false,
    external_vent_as_default: false,
    lc_battery_voltage_volt: 0,
  },
  weather_data: {
    wind_speed_knt: 0,
  },
}
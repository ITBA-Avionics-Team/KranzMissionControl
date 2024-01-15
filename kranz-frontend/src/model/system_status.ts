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

interface SystemStatus {
  on_board: OnBoardSystemStatus;
  launchpad: LaunchpadSystemStatus;
  weather_data: WeatherData;
}
// TODO: UPDATE THESE INTERFACES SO TAHT WE CAN USE THEM

interface FlightComputersStatus {
  altium_ok: boolean;
  ada_ok: boolean;
}

interface OnBoardSystemStatus {
  connection_status: string;
  tank_pressure_bar: number;
  tank_temp_celsius: number;
  tank_depress_vent_temp_celsius: number;
  tank_depress_vent_valve_open: boolean;
  engine_valve_open: boolean;
  combustion_chamber_pressure_bar: number;
  combustion_chamber_temp_celsius: number;
  obec_battery_voltage_volt: number;
  flight_computers_status: FlightComputersStatus;
}

type LCState = 'STANDBY' | 'STANDBY_PRESSURE_WARNING' | 'STANDBY_PRESSURE_WARNING_EXTERNAL_VENT' | 'LOADING' | 'PRE_FLIGHT_CHECK' | 'PRE_LAUNCH_WIND_CHECK' | 'PRE_LAUNCH_UMBRILICAL_DISCONNECT' | 'IGNITION_IGNITERS_ON' | 'IGNITION_OPEN_VALVE' | 'IGNITION_IGNITERS_OFF' | 'ABORT';

interface LaunchpadSystemStatus {
  current_state: LCState;
  connection_status: string;
  loading_line_pressure_bar: number;
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

// export const DefaultSystemStatus: SystemStatus = {
//   on_board: {
//     connection_status: "OK",
//     tank_pressure_bar: 14,
//     tank_temp_celsius: 26,
//     tank_depress_vent_temp_celsius: 10,
//     tank_depress_vent_valve_open: true,
//     engine_valve_open: false,
//     obec_battery_voltage_volt: 0,
//     combustion_chamber_pressure_bar: 20,
//     combustion_chamber_temp_celsius: 25,
//     flight_computers_status: {
//       altium_ok: false,
//       ada_ok: false,
//     },
//   },
//   launchpad: {
//     current_state: 'STANDBY',
//     connection_status: "OK",
//     loading_line_pressure_bar: 0,
//     loading_valve_open: false,
//     loading_depress_vent_valve_open: false,
//     umbrilical_connected: false,
//     igniter_continuity_ok: false,
//     external_vent_as_default: false,
//     lc_battery_voltage_volt: 0,
//   },
//   weather_data: {
//     wind_speed_knt: 0,
//   },
// }

export const DefaultSystemStatus = {
  on_board: {
    connection_status: "?",
    tank_depress_vent_temp_celsius: "?",
    tank_depress_vent_valve_open: "?",
    engine_valve_open: "?",
    obec_battery_voltage_volt: "?",
    flight_computers_status: {
      altium_ok: "?",
      ada_ok: "?",
    },
  },
  launchpad: {
    connection_status: "?",
    current_state: "?",
    loading_line_pressure_bar: "?",
    ground_temp_celsius: "?",
    loading_valve_open: "?",
    loading_depress_vent_valve_open: "?",
    umbrilical_connected: "?",
    umbrilical_finished_disconnect: "?",
    igniter_continuity_ok: "?",
    external_vent_as_default: "?"
  },
  weather_data: {
    wind_speed_knt: "?",
  },
}
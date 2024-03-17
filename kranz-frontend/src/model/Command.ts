export interface Command {
  command_type: CommandType,
  uint_value: number,
  valve: string,
  state: string,
  bool_value: boolean,
  string_value: string
};

const COMMAND_TYPES = ['VALVE_COMMAND', 'SWITCH_STATE_COMMAND', 'SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND', 'EMPTY_COMMAND', 'RAW_COMMAND'] as const;
export type CommandType = typeof COMMAND_TYPES[number];

const STATES = ['STANDBY', 'STANDBY_PRESSURE_WARNING', 'STANDBY_PRESSURE_WARNING_EXTERNAL_VENT', 'LOADING', 'PRE_FLIGHT_CHECK', 'PRE_LAUNCH_WIND_CHECK', 'PRE_LAUNCH_UMBRILICAL_DISCONNECT', 'IGNITION_IGNITERS_ON', 'IGNITION_OPEN_VALVE', 'IGNITION_IGNITERS_OFF', 'ABORT'] as const;
export type LCState = typeof STATES[number];

const VALVES = ['TANK_DEPRESS_VENT_VALVE', 'ENGINE_VALVE', 'LOADING_VALVE', 'LOADING_LINE_DEPRESS_VENT_VALVE'] as const;
export type Valve = typeof VALVES[number];

export function commandToMessage(command: Command) {
  switch (command.command_type) {
    case 'VALVE_COMMAND':
      switch (command.valve) {
        case 'TANK_DEPRESS_VENT_VALVE':
          return `VCTDVV${command.uint_value}|`;
        case 'ENGINE_VALVE':
          return `VCENGV${command.uint_value}|`;
        case 'LOADING_VALVE':
          return `VCLDGV${command.uint_value}|`;
        case 'LOADING_LINE_DEPRESS_VENT_VALVE':
          return `VCLDVV${command.uint_value}|`;
      }
      break;
    case 'SWITCH_STATE_COMMAND':
      return `SS${get4ByteStringFromState(command.state)}|`;
      break;
    case 'SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND':
      return `EV${command.bool_value ? '1' : '0'}|`;
      break;
    case 'EMPTY_COMMAND':
      break;
    case 'RAW_COMMAND':
      return command.string_value;
      break;
  }

}

function get4ByteStringFromState(state) {
  switch (state) {
    case 'STANDBY':
      return 'STBY';
    case 'STANDBY_PRESSURE_WARNING':
      return 'STPW';
    case 'STANDBY_PRESSURE_WARNING_EXTERNAL_VENT':
      return 'STPE';
    case 'LOADING':
      return 'LDNG';
    case 'PRE_FLIGHT_CHECK':
      return 'PRFC';
    case 'PRE_LAUNCH_WIND_CHECK':
      return 'PRLW';
    case 'PRE_LAUNCH_UMBRILICAL_DISCONNECT':
      return 'PRLU';
    case 'IGNITION_IGNITERS_ON':
      return 'IGON';
    case 'IGNITION_OPEN_VALVE':
      return 'IGVO';
    case 'IGNITION_IGNITERS_OFF':
      return 'IGOF';
    case 'ABORT':
      return 'ABRT';
  }
  return '';
}

function isCompleteCommand(command: Command): boolean {
  if (!command.command_type ||
      (command.command_type == 'SWITCH_STATE_COMMAND' && command.state == null) ||
      (command.command_type == 'VALVE_COMMAND' && command.valve == null) ||
      (command.command_type == 'SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND' && command.bool_value == null) ||
      (command.command_type == 'RAW_COMMAND' && command.string_value == null))
      return false;
  return true;
}
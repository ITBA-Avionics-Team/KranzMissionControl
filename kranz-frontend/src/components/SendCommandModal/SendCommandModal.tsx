import React, { useContext, useState } from 'react';
import { SendCommandModalContext } from '../../contexts/SendCommandModalContext';
import CustomCommandSender from '../CustomCommandSender/CustomCommandSender';
import LoadingIndicator from '../LoadingIndicator/LoadingIndicator';
import './SendCommandModal.css'; // Make sure to create this CSS file

const COMMAND_TYPES = ["VALVE_COMMAND","SWITCH_STATE_COMMAND","SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND","RAW_COMMAND"] as const;
export type CommandType = typeof COMMAND_TYPES[number];

const STATES = ["STANDBY","STANDBY_PRESSURE_WARNING","STANDBY_PRESSURE_WARNING_EXTERNAL_VENT","LOADING","PRE_FLIGHT_CHECK","PRE_LAUNCH_WIND_CHECK","PRE_LAUNCH_UMBRILICAL_DISCONNECT","IGNITION_IGNITERS_ON","IGNITION_OPEN_VALVE","IGNITION_IGNITERS_OFF","ABORT"] as const;
export type LCState = typeof STATES[number];

const VALVES = ["TANK_DEPRESS_VENT_VALVE", "ENGINE_VALVE", "LOADING_VALVE", "LOADING_LINE_DEPRESS_VENT_VALVE"] as const;
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
      (command.command_type == 'VALVE_COMMAND' && (command.valve == null || command.uint_value == null)) ||
      (command.command_type == 'SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND' && command.bool_value == null) ||
      (command.command_type == 'RAW_COMMAND' && command.string_value == null))
      return false;
  return true;
}

const SendCommandModal = ({presetCommand}) => {
  const { showSendCommandModal, setShowSendCommandModal } = useContext<boolean>(SendCommandModalContext);
  const [currentCommand, setCurrentCommand] = useState({});
  const [currentCommandRaw, setCurrentCommandRaw] = useState('');
  const [commandStatus, setCommandStatus] = useState('');

  const closeModal = () => setShowSendCommandModal(false);

  const setCurrentCommandAndUpdateCommandRaw = (command) => {
    setCurrentCommand(command);
    if (isCompleteCommand(command)) {
      setCurrentCommandRaw(commandToMessage(command));
    } else {
      setCurrentCommandRaw('');
    }
    setCommandStatus('');
  }

  const setCurrentCommandType = (type) => {
    let updatedCommand = { ...currentCommand, command_type: type };
    setCurrentCommandAndUpdateCommandRaw(updatedCommand);
  };

  const setCurrentCommandState = (state) => {
    let updatedCommand = { ...currentCommand, state: state };
    setCurrentCommandAndUpdateCommandRaw(updatedCommand);
  };

  const setCurrentCommandValve = (valve) => {
    let updatedCommand = { ...currentCommand, valve: valve };
    setCurrentCommandAndUpdateCommandRaw(updatedCommand);
  };

  const setCurrentCommandUintValue = (value) => {
    let updatedCommand = { ...currentCommand, uint_value: value };
    setCurrentCommandAndUpdateCommandRaw(updatedCommand);
  };

  const setCurrentCommandBoolValue = (value) => {
    let updatedCommand = { ...currentCommand, bool_value: value };
    setCurrentCommandAndUpdateCommandRaw(updatedCommand);
  };

  const sendCommand = () => {
    const data = {
      command_type: "RAW_COMMAND",
      string_value: currentCommandRaw
    };

    setCommandStatus('Sending command...');

    fetch('http://localhost:8080/command', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json'
      },
      body: JSON.stringify(data)
    })
    .then(response => {
      if (response.ok) {
        return response.json();
      }
      throw new Error('Network response was not ok.');
    })
    .then(data =>{
      setCommandStatus('Command sent successfully');
    })
    .catch(error => {
      setCommandStatus('There was an error sending the command:' + error);
    });
  };

  return (
    <>
      {showSendCommandModal && (
        <div className="modal">
          <div className="modal-content">
            <span className="close-button" onClick={closeModal}>&times;</span>
            <h2>Send Command</h2>
            {/* <CommandTypeSelector /> */}

            <h3>Command type</h3>
            {COMMAND_TYPES.map((type) => (
                <button
                  key={type}
                  className={currentCommand.command_type === type ? 'selected' : ''}
                  onClick={() => setCurrentCommandType(type)}>
                  {type}
                </button>
              ))}

            {currentCommand.command_type == 'SWITCH_STATE_COMMAND' && (
              <>
              <h3>State</h3>
              {STATES.map((state) => (
                <button
                  key={state}
                  className={currentCommand.state === state ? 'selected' : ''}
                  onClick={() => setCurrentCommandState(state)}>
                  {state}
                </button>
              ))}
              </>
            )
            }

            {currentCommand.command_type == 'VALVE_COMMAND' && (
              <>
              <h3>Valve</h3>
              {VALVES.map((valve) => (
                <button
                  key={valve}
                  className={currentCommand.valve === valve ? 'selected' : ''}
                  onClick={() => setCurrentCommandValve(valve)}>
                  {valve}
                </button>
              ))}

              <h3>Action</h3>
              <button
                  className={currentCommand.uint_value === 1 ? 'selected' : ''}
                  onClick={() => setCurrentCommandUintValue(1)}>
                  Open
              </button>
              <button
                  className={currentCommand.uint_value === 0 ? 'selected' : ''}
                  onClick={() => setCurrentCommandUintValue(0)}>
                  Close
              </button>
              </>
            )
            }

            {currentCommand.command_type == 'SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND' && (
              <>
              <h3>External vent as default</h3>
              <button
                  className={currentCommand.bool_value === true ? 'selected' : ''}
                  onClick={() => setCurrentCommandBoolValue(true)}>
                  True
              </button>
              <button
                  className={currentCommand.bool_value === false ? 'selected' : ''}
                  onClick={() => setCurrentCommandBoolValue(false)}>
                  False
              </button>
              </>
            )
            }
            
            <div>
              <input
                type="text"
                value={currentCommandRaw}
                onChange={(event)=>setCurrentCommandRaw(event.target.value)}
                placeholder="Enter command"
              />
              <button onClick={sendCommand}>Send Command</button>
              <div>{commandStatus}</div>
              {commandStatus == 'Sending command...' && (<LoadingIndicator />)}
            </div>
          </div>
        </div>
      )}
    </>
  );
};

export default SendCommandModal;

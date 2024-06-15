import React, { useContext, useEffect, useState } from 'react';
import { SendCommandModalContext } from '../../contexts/SendCommandModalContext';
import CustomCommandSender from '../CustomCommandSender/CustomCommandSender';
import LoadingIndicator from '../LoadingIndicator/LoadingIndicator';
import './SendCommandModal.css'; // Make sure to create this CSS file

const COMMAND_TYPES = ["VALVE_COMMAND","SWITCH_STATE_COMMAND","SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND","RAW_COMMAND"] as const;
export type CommandType = typeof COMMAND_TYPES[number];

const STATES = ["STANDBY","STANDBY_PRESSURE_WARNING","LOADING","PRE_FLIGHT_CHECK","PRE_LAUNCH_WIND_CHECK","PRE_LAUNCH_UMBRILICAL_DISCONNECT","IGNITION_OPEN_VALVE", "IGNITION_IGNITERS_ON","IGNITION_IGNITERS_OFF","ABORT"] as const;
export type LCState = typeof STATES[number];

const VALVES = ["ENGINE_VALVE", "LOADING_VALVE"] as const;
export type Valve = typeof VALVES[number];

export function commandToMessage(command: Command) {
  switch (command.type) {
    case 'VALVE_COMMAND':
      switch (command.valve) {
        case 'TANK_DEPRESS_VENT_VALVE':
          return `VCTDVV${command.uintValue}|`;
        case 'ENGINE_VALVE':
          return `VCENGV${command.uintValue}|`;
        case 'LOADING_VALVE':
          return `VCLDGV${command.uintValue}|`;
        case 'LOADING_LINE_DEPRESS_VENT_VALVE':
          return `VCLDVV${command.uintValue}|`;
      }
      break;
    case 'SWITCH_STATE_COMMAND':
      return `SS${get4ByteStringFromState(command.state)}|`;
      break;
    case 'SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND':
      return `EV${command.boolValue ? '1' : '0'}|`;
      break;
    case 'EMPTY_COMMAND':
      break;
    case 'RAW_COMMAND':
      return command.stringValue;
      break;
  }

}

function get4ByteStringFromState(state) {
  switch (state) {
    case 'STANDBY':
      return 'STBY';
    case 'STANDBY_PRESSURE_WARNING':
      return 'STPW';
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
  if (!command.type ||
      (command.type == 'SWITCH_STATE_COMMAND' && command.state == null) ||
      (command.type == 'VALVE_COMMAND' && (command.valve == null || command.uintValue == null)) ||
      (command.type == 'SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND' && command.boolValue == null) ||
      (command.type == 'RAW_COMMAND' && command.stringValue == null))
      return false;
  return true;
}

const SendCommandModal = ({presetCommand}) => {
  const { sendCommandModalContext, setSendCommandModalContext } = useContext(SendCommandModalContext);

  const closeModal = () => setSendCommandModalContext({showModal: false, currentCommand: {}});

  const setCurrentCommandAndUpdateCommandRaw = (updatedContext) => {
    if (isCompleteCommand(updatedContext.currentCommand)) {
      setSendCommandModalContext({...updatedContext, currentCommand: {...updatedContext.currentCommand, raw: commandToMessage(updatedContext.currentCommand)}});
    } else {
      setSendCommandModalContext({...updatedContext, currentCommand: {...updatedContext.currentCommand, raw: ''}});
    }
  }

  const setCurrentCommandType = (type) => {
    setCurrentCommandAndUpdateCommandRaw({...sendCommandModalContext, currentCommand: {...sendCommandModalContext.currentCommand, type: type}});
  };

  const setCurrentCommandState = (state) => {
    setCurrentCommandAndUpdateCommandRaw({...sendCommandModalContext, currentCommand: {...sendCommandModalContext.currentCommand, state: state}});
  };

  const setCurrentCommandValve = (valve) => {
    setCurrentCommandAndUpdateCommandRaw({...sendCommandModalContext, currentCommand: {...sendCommandModalContext.currentCommand, valve: valve}});
  };

  const setCurrentCommandUintValue = (value) => {
    setCurrentCommandAndUpdateCommandRaw({...sendCommandModalContext, currentCommand: {...sendCommandModalContext.currentCommand, uintValue: value}});
  };

  const setCurrentCommandBoolValue = (value) => {
    setCurrentCommandAndUpdateCommandRaw({...sendCommandModalContext, currentCommand: {...sendCommandModalContext.currentCommand, boolValue: value}});
  };

  const setCurrentCommandRaw = (value) => {
    setCurrentCommandAndUpdateCommandRaw({...sendCommandModalContext, currentCommand: {...sendCommandModalContext.currentCommand, raw: value}});
  };

  const setCommandStatus = (status) => {
    setCurrentCommandAndUpdateCommandRaw({...sendCommandModalContext, commandStatus: status});
  };

  const sendCommand = () => {
    const data = {
      command_type: "RAW_COMMAND",
      string_value: sendCommandModalContext.currentCommand.raw
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
      {sendCommandModalContext.showModal && (
        <div className="modal">
          <div className="modal-content">
            <span className="close-button" onClick={closeModal}>&times;</span>
            <h2>Send Command</h2>
            {/* <CommandTypeSelector /> */}

            <h3>Command type</h3>
            {COMMAND_TYPES.map((type) => (
                <button
                  key={type}
                  className={sendCommandModalContext.currentCommand.type === type ? 'selected' : ''}
                  onClick={() => setCurrentCommandType(type)}>
                  {type}
                </button>
              ))}

            {sendCommandModalContext.currentCommand.type == 'SWITCH_STATE_COMMAND' && (
              <>
              <h3>State</h3>
              {STATES.map((state) => (
                <button
                  key={state}
                  className={sendCommandModalContext.currentCommand.state === state ? 'selected' : ''}
                  onClick={() => setCurrentCommandState(state)}>
                  {state}
                </button>
              ))}
              </>
            )
            }

            {sendCommandModalContext.currentCommand.type == 'VALVE_COMMAND' && (
              <>
              <h3>Valve</h3>
              {VALVES.map((valve) => (
                <button
                  key={valve}
                  className={sendCommandModalContext.currentCommand.valve === valve ? 'selected' : ''}
                  onClick={() => setCurrentCommandValve(valve)}>
                  {valve}
                </button>
              ))}

              <h3>Action</h3>
              <button
                  className={sendCommandModalContext.currentCommand.uintValue === 1 ? 'selected' : ''}
                  onClick={() => setCurrentCommandUintValue(1)}>
                  Open
              </button>
              <button
                  className={sendCommandModalContext.currentCommand.uintValue === 0 ? 'selected' : ''}
                  onClick={() => setCurrentCommandUintValue(0)}>
                  Close
              </button>
              </>
            )
            }

            {sendCommandModalContext.currentCommand.type == 'SET_EXTERNAL_VENT_AS_DEFAULT_COMMAND' && (
              <>
              <h3>External vent as default</h3>
              <button
                  className={sendCommandModalContext.currentCommand.boolValue === true ? 'selected' : ''}
                  onClick={() => setCurrentCommandBoolValue(true)}>
                  True
              </button>
              <button
                  className={sendCommandModalContext.currentCommand.boolValue === false ? 'selected' : ''}
                  onClick={() => setCurrentCommandBoolValue(false)}>
                  False
              </button>
              </>
            )
            }
            
            <div>
              <input
                type="text"
                value={sendCommandModalContext.currentCommand.raw}
                onChange={(event)=>setCurrentCommandRaw(event.target.value)}
                placeholder="Enter command"
              />
              <button onClick={sendCommand}>Send Command</button>
              <div>{sendCommandModalContext.commandStatus}</div>
              {sendCommandModalContext.commandStatus == 'Sending command...' && (<LoadingIndicator />)}
            </div>
          </div>
        </div>
      )}
    </>
  );
};

export default SendCommandModal;

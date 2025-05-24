import type { JSX } from 'react';
import React, { useContext, ChangeEvent } from 'react';
import { SendCommandModalContext } from '../../contexts/SendCommandModalContext';
import LoadingIndicator from '../LoadingIndicator/LoadingIndicator';
import './SendCommandModal.css';

const COMMAND_TYPES = ["RAW_COMMAND", "STARTUP_SIGNAL_COMMAND"] as const;
export type CommandType = typeof COMMAND_TYPES[number];

export function commandToMessage(command: { type: CommandType; stringValue?: string }) {
  switch (command.type) {
    case 'RAW_COMMAND':
      return command.stringValue || '';
    case 'STARTUP_SIGNAL_COMMAND':
      return 'STARTUP_SIGNAL|';
  }
}

const SendCommandModal: React.FC = () => {
  const { sendCommandModalContext, setSendCommandModalContext } = useContext(SendCommandModalContext);

  const closeModal = () => setSendCommandModalContext({ showModal: false, currentCommand: {} });

  const setCurrentCommandRaw = (value: string) => {
    setSendCommandModalContext({
      ...sendCommandModalContext,
      currentCommand: { type: 'RAW_COMMAND', stringValue: value, raw: value },
    });
  };

  const setCommandStatus = (status: string) => {
    setSendCommandModalContext({ ...sendCommandModalContext, commandStatus: status });
  };

  const sendCommand = (type: CommandType) => {
    let data;
    if (type === 'RAW_COMMAND') {
      data = {
        command_type: 'RAW_COMMAND',
        string_value: sendCommandModalContext.currentCommand.raw,
      };
    } else if (type === 'STARTUP_SIGNAL_COMMAND') {
      data = {
        command_type: 'STARTUP_SIGNAL_COMMAND'
      };
    }
    setCommandStatus('Sending command...');
    fetch('http://localhost:8080/command', {
      method: 'POST',
      headers: {
        'Content-Type': 'application/json',
      },
      body: JSON.stringify(data),
    })
      .then((response) => {
        if (response.ok) {
          return response.json();
        }
        throw new Error('Network response was not ok.');
      })
      .then(() => {
        setCommandStatus('Command sent successfully');
      })
      .catch((error) => {
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
            <div>
              <input
                type="text"
                value={sendCommandModalContext.currentCommand.raw || ''}
                onChange={(event: ChangeEvent<HTMLInputElement>) => setCurrentCommandRaw(event.target.value)}
                placeholder="Enter raw command"
              />
              <button onClick={() => sendCommand('RAW_COMMAND')}>Send Raw Command</button>
            </div>
            <div style={{ marginTop: '1em' }}>
              <button onClick={() => sendCommand('STARTUP_SIGNAL_COMMAND')}>Send STARTUP SIGNAL</button>
            </div>
            <div>{sendCommandModalContext.commandStatus}</div>
            {sendCommandModalContext.commandStatus === 'Sending command...' && <LoadingIndicator />}
          </div>
        </div>
      )}
    </>
  );
};

export default SendCommandModal;

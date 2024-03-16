import React, { useState } from 'react';

function CommandSender() {
  const [command, setCommand] = useState('');

  const handleInputChange = (event) => {
    setCommand(event.target.value);
  };

  const sendCommand = () => {
    const data = {
      command_type: "RAW_COMMAND",
      string_value: command
    };

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
    .then(data => console.log('Command sent successfully:', data))
    .catch(error => console.error('There was an error sending the command:', error));
  };

  return (
    <div>
      <input
        type="text"
        value={command}
        onChange={handleInputChange}
        placeholder="Enter command"
      />
      <button onClick={sendCommand}>Send Command</button>
    </div>
  );
}

export default CommandSender;

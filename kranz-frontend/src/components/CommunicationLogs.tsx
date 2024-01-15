import React, { useEffect, useState } from 'react';

const CommunicationLogs = () => {
  const [messages, setMessages] = useState([]);

  useEffect(() => {
    const ws = new WebSocket('ws://127.0.0.1:8080/system_status');

    ws.onmessage = (event) => {
      setMessages((prevMessages) => [...prevMessages, event.data]);
    };

    ws.onerror = (error) => {
      // Handle the error
    };

    return () => {
      ws.close();
    };
  }, []);

  return (
    <div>
      {messages.map((message, index) => (
        <p key={index}>{message}</p>
      ))}
    </div>
  );
};

export default CommunicationLogs;

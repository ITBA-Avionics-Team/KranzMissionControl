import React, { useEffect, useState } from 'react';

{

}

const CommunicationLogs = () => {
  const [latestSystemStatus, setLatestSystemStatus] = useState();

  useEffect(() => {
    const ws = new WebSocket('ws://127.0.0.1:8080/system_status');

    ws.onmessage = (event) => {
      setLatestSystemStatus((previousSystemStatus) => {
        return JSON.parse(event.data)
      });
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
      {latestSystemStatus?.launchpad?.current_state}
    </div>
  );
};

export default CommunicationLogs;

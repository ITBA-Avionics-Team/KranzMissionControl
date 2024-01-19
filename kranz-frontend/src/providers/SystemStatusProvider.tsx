import { useEffect, useState } from "react";
import { SystemStatusContext } from "../contexts/SystemStatusContext";
import { DefaultSystemStatus } from "../model/SystemStatus";

export const SystemStatusProvider = ({ children }) => {
  const [latestSystemStatus, setLatestSystemStatus] = useState(DefaultSystemStatus);


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
    <SystemStatusContext.Provider value={{ latestSystemStatus, setLatestSystemStatus }}>
      {children}
    </SystemStatusContext.Provider>
  );
};

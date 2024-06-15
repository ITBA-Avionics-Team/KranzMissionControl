import { useEffect, useState } from "react";
import { SystemStatusContext } from "../contexts/SystemStatusContext";
import { DefaultSystemStatus } from "../model/SystemStatus";

export const SystemStatusProvider = ({ children }) => {
  const [latestSystemStatus, setLatestSystemStatus] = useState(DefaultSystemStatus);

  let latestStatusUpdateTime = new Date();


  useEffect(() => {
    const ws = new WebSocket('ws://127.0.0.1:8080/system_status');

    ws.onmessage = (event) => {
      setLatestSystemStatus((previousSystemStatus) => {
        let systemStatusJsonData = JSON.parse(event.data);
        latestStatusUpdateTime = new Date();
        return systemStatusJsonData;
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

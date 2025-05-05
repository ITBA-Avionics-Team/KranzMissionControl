import { useContext } from "react";
import { SystemStatusContext } from "../contexts/SystemStatusContext";
import { SystemStatus } from "../model/SystemStatus";
import './SystemVisualizer.css'
import SystemStatusText from "./SystemStatusText";
import LoadingDiagram from "./LoadingDiagram/LoadingDiagram";
import StatusBar from "./StatusBar";
import FlightDiagram from "./FlightDiagram/FlightDiagram";

const SystemVisualizer = () => {
  const { latestSystemStatus, setLatestSystemStatus } = useContext(SystemStatusContext);

  return (
    <div className="system-visualizer">
      <div className="system-visualizer-content">
        {/* <LoadingDiagram systemStatus={latestSystemStatus}/> */}
        <FlightDiagram systemStatus={latestSystemStatus} />
        <SystemStatusText systemStatus={latestSystemStatus} />
      </div>
      <StatusBar systemStatus={latestSystemStatus} />
    </div>
    
  );
};

export default SystemVisualizer;
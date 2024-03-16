import { useContext } from "react";
import { SystemStatusContext } from "../contexts/SystemStatusContext";
import { SystemStatus } from "../model/SystemStatus";
import './SystemVisualizer.css'
import SystemStatusText from "./SystemStatusText";
import LoadingDiagram from "./LoadingDiagram/LoadingDiagram";
import StatusBar from "./StatusBar";

const SystemVisualizer = () => {
  const { latestSystemStatus, setLatestSystemStatus } = useContext<SystemStatus>(SystemStatusContext);

  return (
    <div className="system-visualizer">
      <div className="system-visualizer-content">
        <LoadingDiagram systemStatus={latestSystemStatus}/>
        <SystemStatusText systemStatus={latestSystemStatus} />
      </div>
      <StatusBar systemStatus={latestSystemStatus} />
    </div>
    
  );
};

export default SystemVisualizer;
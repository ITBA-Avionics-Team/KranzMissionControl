import { useContext } from "react";
import { SystemStatusContext } from "../contexts/SystemStatusContext";
import { SystemStatus } from "../model/SystemStatus";
import LoadingDiagram from "./LoadingDiagram/LoadingDiagram";
import StatusBar from "./StatusBar";

const SystemVisualizer = () => {
  const { latestSystemStatus, setLatestSystemStatus } = useContext<SystemStatus>(SystemStatusContext);

  return (
    <>
      <LoadingDiagram systemStatus={latestSystemStatus}/>
      <StatusBar systemStatus={latestSystemStatus} />
    </>
  );
};

export default SystemVisualizer;
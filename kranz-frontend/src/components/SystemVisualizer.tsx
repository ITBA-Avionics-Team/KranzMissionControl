import { useContext } from "react";
import { SystemStatusContext } from "../contexts/SystemStatusContext";
import { SystemStatus } from "../model/SystemStatus";
import LoadingDiagram from "./LoadingDiagram/LoadingDiagram";

const SystemVisualizer = () => {
  const { latestSystemStatus, setLatestSystemStatus } = useContext<SystemStatus>(SystemStatusContext);

  return (
    <div>
      {/* {latestSystemStatus} */}
      <h3>On Board Engine Computer</h3>
      <div>Tank depress vent valve open: {latestSystemStatus?.on_board?.tank_depress_vent_valve_open ? "True" : "False"}</div>
      <div>Engine valve open: {latestSystemStatus?.on_board?.engine_valve_open  ? "True" : "False"}</div>
      <LoadingDiagram systemStatus={latestSystemStatus}/>
    </div>
  );
};

export default SystemVisualizer;
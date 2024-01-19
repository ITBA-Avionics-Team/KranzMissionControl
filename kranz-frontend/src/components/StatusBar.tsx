const StatusBar = (systemStatus) => {

  const statusBarStyle = {
    width: '100vw',
    height: '5rem',
    backgroundColor: 'grey'
  }

  return (
    <div style={statusBarStyle}>
      <h3>On Board Engine Computer</h3>
      <div>Tank depress vent valve open: {systemStatus?.on_board?.tank_depress_vent_valve_open ? "True" : "False"}</div>
      <div>Engine valve open: {systemStatus?.on_board?.engine_valve_open  ? "True" : "False"}</div>
    </div>
  );
};

export default StatusBar;
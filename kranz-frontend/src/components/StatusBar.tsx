const StatusBar = ({systemStatus}) => {

  // console.log(systemStatus)

  const statusBarStyle = {
    width: '100vw',
    height: '5rem',
    backgroundColor: 'var(--main-bg-color)',
    color: 'var(--main-fg-color)',

    borderColor: 'var(--main-fg-color)',
    borderTopStyle: 'solid',
    borderWidth: '1px',

    display: 'flex',
    flexDirection: 'row',
    justifyContent: 'space-between',

    // padding: '0 2rem'
  }

  const statusBarElemStyle = {
    display: 'flex',
    flexDirection: 'column',
    margin: '0 3rem'
  }

  return (
    <div style={statusBarStyle}>
      <div style={statusBarElemStyle}>
        <div><b>LC State:</b> {systemStatus.launchpad.current_state}</div>
        <div><b>LC Connection status:</b> {systemStatus.launchpad.connection_status}</div>
      </div>
      <div style={statusBarElemStyle}>
        {/* <div><b>OBEC State:</b> {systemStatus.on_board?.current_state}</div> */}
        <div><b>OBEC Connection status:</b> {systemStatus.on_board?.connection_status}</div>
      </div>

      <div style={statusBarElemStyle}>
        <div><b>Default venting mode:</b> {systemStatus.launchpad?.external_vent_as_default ? "EXTERNAL" : "INTERNAL"}</div>
      </div>

    </div>
  );
};

export default StatusBar;
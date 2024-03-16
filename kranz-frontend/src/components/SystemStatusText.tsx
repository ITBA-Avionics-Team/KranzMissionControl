import React, { useContext, useEffect, useState } from 'react';
import { SystemStatusContext } from '../contexts/SystemStatusContext';
import { SystemStatus } from '../model/SystemStatus';
import './SystemStatusText.css'

const SystemStatusText = ({ systemStatus }) => {

  return (
    <div className="system-status-text">
      <div><b>Current state:</b> {systemStatus?.launchpad?.current_state}</div>
      <div><b>Tank pressure:</b> {systemStatus?.on_board?.tank_pressure_psi} PSI</div>
      <div><b>Tank temperature:</b> {systemStatus?.on_board?.tank_temp_celsius} °C</div>
      <div><b>Tank depress vent temperature:</b> {systemStatus?.on_board?.tank_depress_vent_temp_celsius} °C</div>
      <div><b>Loading line pressure:</b> {systemStatus?.launchpad?.load_line_pressure_psi} PSI</div>
      <div><b>OBEC battery voltage:</b> {systemStatus?.on_board?.obec_battery_voltage_volt} V</div>
      <div><b>LC battery voltage:</b> {systemStatus?.launchpad?.lc_battery_voltage_volt} V</div>
      <div><b>Connection to OBEC:</b> {systemStatus?.on_board?.connection_status ? "OK" : "Error"}</div>
      <div><b>Tank depress vent valve:</b> {systemStatus?.on_board?.tank_depress_vent_valve_open ? "Open" : "Closed"}</div>
      <div><b>Engine valve:</b> {systemStatus?.on_board?.engine_valve_open ? "Open" : "Closed"}</div>
      <div><b>Loading Valve:</b> {systemStatus?.launchpad?.loading_valve_open ? "Open" : "Closed"}</div>
      <div><b>Loading depress vent valve:</b> {systemStatus?.launchpad?.loading_depress_vent_valve_open ? "Open" : "Closed"}</div>
      <div><b>Hydraulic umbrilical:</b> {systemStatus?.launchpad?.umbrilical_connected ? "Connected" : "Disconnected"}</div>
      <div><b>Igniter continuity:</b> {systemStatus?.launchpad?.igniter_continuity_ok ? "Ok" : "Error"}</div>
      <div><b>External vent as default:</b> {systemStatus?.launchpad?.external_vent_as_default ? "True" : "False"}</div>
      <div><b>Wind:</b> {systemStatus?.weather_data?.wind_speed_knt}kt</div>
    </div>
  );
};

export default SystemStatusText;

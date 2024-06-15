import React, { useContext, useEffect, useState } from 'react';
import { SystemStatusContext } from '../contexts/SystemStatusContext';
import { SystemStatus } from '../model/SystemStatus';
import './SystemStatusText.css'

const SystemStatusText = ({ systemStatus }) => {

  return (
    <div className="system-status-text">
      <table>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Current state:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.launchpad?.current_state}</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Tank depress vent temperature:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.on_board?.tank_depress_vent_temp_celsius} °C</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Loading line pressure:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.launchpad?.loading_line_pressure_bar} BAR</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Ground temperature:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.launchpad?.ground_temp_celsius} BAR</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>OBEC battery voltage:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.on_board?.obec_battery_voltage_volt} V</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Connection to OBEC:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.on_board?.connection_status}</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Engine valve:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.on_board?.engine_valve_open ? "Open" : "Closed"}</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Loading Valve:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.launchpad?.loading_valve_open ? "Open" : "Closed"}</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Loading depress vent valve:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.launchpad?.loading_depress_vent_valve_open ? "Open" : "Closed"}</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Hydraulic umbrilical connected:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.launchpad?.umbrilical_connected ? "Connected" : "Disconnected"}</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Hydraulic umbrilical finished disconnect:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.launchpad?.umbrilical_finished_disconnect ? "Finished" : "Still disconnecting"}</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Igniter continuity:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.launchpad?.igniter_continuity_ok ? "Ok" : "Error"}</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>External vent as default:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.launchpad?.external_vent_as_default ? "True" : "False"}</td>
        </tr>
        <tr>
          <td colSpan="2" style={{ fontWeight: 'bold' }}>Wind:</td>
          <td colSpan="2" style={{ textAlign: 'right' }}> {systemStatus?.weather_data?.wind_speed_knt}kt</td>
        </tr>
      </table>
    </div>
  );
};

export default SystemStatusText;

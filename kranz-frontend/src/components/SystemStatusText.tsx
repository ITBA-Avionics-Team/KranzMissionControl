import React from 'react';
import { SystemStatus } from '../model/SystemStatus';
import './SystemStatusText.css'

interface SystemStatusTextProps {
  systemStatus: SystemStatus;
}

const SystemStatusText: React.FC<SystemStatusTextProps> = ({ systemStatus }) => {
  return (
    <div className="system-status-text">
      <table>
        <tbody>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>Flight Status:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.status || '?'}</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>Altitude:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.altitude || '?'} m</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>Tilt:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.tilt || '?'}°</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>GPS Latitude:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.gps_latitude || '?'}</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>GPS Longitude:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.gps_longitude || '?'}</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>GPS Altitude:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.gps_altitude || '?'} m</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>Acceleration:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.acceleration || '?'} m/s²</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>Temperature:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.temperature || '?'} °C</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default SystemStatusText;

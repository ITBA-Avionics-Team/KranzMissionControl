import * as React from 'react';
import { SystemStatus } from '../model/SystemStatus';
import './SystemStatusText.css';

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
            <td colSpan={2} style={{ fontWeight: 'bold' }}>Mission Time:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.mission_time || '?'}</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>Packet Count:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.packet_count || '?'}</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>Battery Level:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.battery_level || '?'} %</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>IMU Y Velocity:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.imu_y_vel || '?'} m/s</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>IMU Roll:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.imu_roll || '?'}°</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>IMU Pitch:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.imu_pitch || '?'}°</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>GNSS Time:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.gnss_time || '?'}</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>GNSS Latitude:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.gnss_latitude || '?'}</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>GNSS Longitude:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.gnss_longitude || '?'}</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>GNSS Altitude:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.gnss_altitude || '?'} m</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>BME Pressure:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.bme_pressure || '?'} Pa</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>BME Altitude:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.bme_altitude || '?'} m</td>
          </tr>
          <tr>
            <td colSpan={2} style={{ fontWeight: 'bold' }}>BME Temperature:</td>
            <td colSpan={2} style={{ textAlign: 'right' }}>{systemStatus?.flight_telemetry?.bme_temperature || '?'} °C</td>
          </tr>
        </tbody>
      </table>
    </div>
  );
};

export default SystemStatusText;

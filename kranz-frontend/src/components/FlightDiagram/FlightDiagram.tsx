import React from 'react';
import { SystemStatus } from '../../model/SystemStatus';
import './FlightDiagram.css';

interface FlightDiagramProps {
  systemStatus: SystemStatus;
}

const FlightDiagram: React.FC<FlightDiagramProps> = ({ systemStatus }) => {
  // Get GPS coordinates from flight telemetry if available
  const latitude = systemStatus?.flight_telemetry?.gps_latitude;
  const longitude = systemStatus?.flight_telemetry?.gps_longitude;
  const altitude = systemStatus?.flight_telemetry?.gps_altitude;
  
  // Check if we have valid coordinates
  const hasValidCoordinates = 
    typeof latitude === 'number' && 
    typeof longitude === 'number' && 
    !isNaN(latitude) && 
    !isNaN(longitude);
  
  // Default position (Spaceport America)
  const defaultLat = 32.9895;
  const defaultLng = -106.9749;
  
  // Use valid coordinates or default
  const mapLat = hasValidCoordinates ? latitude : defaultLat;
  const mapLng = hasValidCoordinates ? longitude : defaultLng;
  
  // Map parameters
  const zoom = 15;
  const mapSize = 600;
  const markerSize = 16;
  
  // Generate static map URL
  const staticMapUrl = `https://www.openstreetmap.org/export/embed.html?bbox=${mapLng-0.01},${mapLat-0.01},${mapLng+0.01},${mapLat+0.01}&layer=mapnik&marker=${mapLat},${mapLng}`;
  
  return (
    <div className="flight-diagram">
      <h3>Flight Position</h3>
      <div className="map-container">
        <iframe 
          width="100%" 
          height="300" 
          frameBorder="0" 
          scrolling="no" 
          marginHeight={0} 
          marginWidth={0} 
          src={staticMapUrl}
        ></iframe>
      </div>
      <div className="coordinate-display">
        <p>Latitude: {hasValidCoordinates ? latitude.toFixed(6) : '?'}</p>
        <p>Longitude: {hasValidCoordinates ? longitude.toFixed(6) : '?'}</p>
        <p>Altitude: {altitude || '?'} m</p>
      </div>
    </div>
  );
};

export default FlightDiagram; 
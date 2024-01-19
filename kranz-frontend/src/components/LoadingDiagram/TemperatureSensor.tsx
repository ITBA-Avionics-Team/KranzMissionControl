import React from 'react';

const TemperatureSensor = ({ name, coords, onHover }) => {

	const xPercentage = coords.x; // 20% from the left edge
  const yPercentage = coords.y; // 30% from the top edge

  // Style for the valve
  const sensorStyle = {
    fill: 'grey',
    stroke: 'black',
    strokeWidth: 2,

		position: 'absolute',
    top: `${yPercentage}%`,
    left: `${xPercentage}%`,
  };

  return (
    <svg style={sensorStyle} onHover={() => onHover(name)}>
      <circle cx="50" cy="50" r="30"/>
    </svg>
  );
};

export default TemperatureSensor;
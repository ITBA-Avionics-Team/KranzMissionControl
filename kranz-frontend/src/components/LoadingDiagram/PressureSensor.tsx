import React from 'react';

const PresssureSensor = ({ name, coords, onHover }) => {

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
    <svg style={sensorStyle}>
      <circle cx="50" cy="50" r="20" onHover={() => onHover(name)}/>
    </svg>
  );
};

export default PresssureSensor;
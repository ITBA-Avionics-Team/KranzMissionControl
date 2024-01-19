import React from 'react';

const Valve = ({ name, coords, status, onClick, onHover }) => {
  // Define the valve's appearance based on its state (open or closed)
  const valveColor = status ? 'green' : 'red';

	const xPercentage = coords.x; // 20% from the left edge
  const yPercentage = coords.y; // 30% from the top edge

  // Style for the valve
  const valveStyle = {
    fill: valveColor,
    stroke: 'black',
    strokeWidth: 2,
    cursor: 'pointer',

		position: 'absolute',
    top: `${yPercentage}%`,
    left: `${xPercentage}%`,
  };

  return (
    <svg style={valveStyle}>
      <circle cx="50" cy="50" r="40" onClick={() => onClick(name)} onHover={() => onHover(name)}/>
    </svg>
  );
};

export default Valve;
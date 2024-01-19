import React from 'react';

const Umbrilical = ({ name, coords, state, onClick, onHover }) => {
  const valveColor = state === 'open' ? 'green' : 'red';

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
    <svg style={valveStyle} >
      <circle cx="50" cy="50" r="40" onClick={() => onClick(name)} onHover={() => onHover(name)}/>
    </svg>
  );
};

export default Umbrilical;
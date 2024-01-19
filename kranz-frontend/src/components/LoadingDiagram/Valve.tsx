import React from 'react';

const Valve = ({ name, coords, status, onClick, onMouseMove, onMouseEnter, onMouseLeave }) => {
  // Define the valve's appearance based on its state (open or closed)
  const valveColor = status ? 'limegreen' : 'red';

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
      <svg onClick={() => onClick(name)} onMouseMove={onMouseMove} onMouseEnter={() => onMouseEnter(name)} onMouseLeave={() => onMouseLeave(name)} style={valveStyle} xmlns="http://www.w3.org/2000/svg" height="24" viewBox="0 -960 960 960" width="24">
        <path d="M440-640v-120H280v-80h400v80H520v120h-80ZM160-120v-320h80v40h120v-120h-40v-80h320v80h-40v120h120v-40h80v320h-80v-40H240v40h-80Zm80-120h480v-80H520v-200h-80v200H240v80Zm240 0Z"/>
      </svg>
  );
};

export default Valve;
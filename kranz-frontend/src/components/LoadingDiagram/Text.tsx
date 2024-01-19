import React from 'react';

const Text = ({ name, value, coords, onMouseEnter, onMouseLeave }) => {

	const xPercentage = coords.x; 
  const yPercentage = coords.y; 

  // Style for the valve
  const textStyle = {
    color: 'var(--main-fg-color)',
    cursor: 'default',

		position: 'absolute',
    top: `${yPercentage}%`,
    left: `${xPercentage}%`,
  };

  return (
    <div style={textStyle} onMouseEnter={() => onMouseEnter(name)} onMouseLeave={() => onMouseLeave(name)}>
      {value}
    </div>
  );
};

export default Text;
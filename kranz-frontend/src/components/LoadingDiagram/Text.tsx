import React from 'react';

const Text = ({ value, coords }) => {

	const xPercentage = coords.x; 
  const yPercentage = coords.y; 

  // Style for the valve
  const textStyle = {
    color: 'black',

		position: 'absolute',
    top: `${yPercentage}%`,
    left: `${xPercentage}%`,
  };

  return (
    <div style={textStyle}>
      {value}
    </div>
  );
};

export default Text;
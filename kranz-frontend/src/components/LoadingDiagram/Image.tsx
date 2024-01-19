import React from 'react';

const Image = ({ src, coords }) => {

	const xPercentage = coords.x; 
  const yPercentage = coords.y; 

  // Style for the valve
  const imgStyle = {
		position: 'absolute',
    top: `${yPercentage}%`,
    left: `${xPercentage}%`,
  };

  return (
    <img src={src} style={imgStyle}/>
  );
};

export default Image;
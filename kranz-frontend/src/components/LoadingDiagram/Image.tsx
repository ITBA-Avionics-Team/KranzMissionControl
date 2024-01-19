import React from 'react';

const Image = ({ src, coords, dimensions }) => {

  // Style for the valve
  const imgStyle = {
		position: 'absolute',
    top: `${coords.y}%`,
    left: `${coords.x}%`,
    width: `${dimensions.x}%`,
    height: `${dimensions.y}%`,
  };

  return (
    <img src={src} style={imgStyle}/>
  );
};

export default Image;
import tankLightModeSvg from '../../assets/loading-diagram/tank_image.svg'
import tankDarkModeSvg from '../../assets/loading-diagram/tank_image_dark_mode.svg'

const Tank = ({ src, coords, dimensions, darkMode }) => {

  // Style for the valve
  const imgStyle = {
		position: 'absolute',
    top: `${coords.y}%`,
    left: `${coords.x}%`,
    // width: `${dimensions.x}%`,
    height: `${dimensions.y}%`,
  };

  return (
    <img src={darkMode ? tankDarkModeSvg : tankLightModeSvg} style={imgStyle}/>
  );
};

export default Tank;
import tankLightModeSvg from '../../assets/loading-diagram/tank_image.svg'
import tankDarkModeSvg from '../../assets/loading-diagram/tank_image_dark_mode.svg'

const Tank = ({ src, coords, dimensions, darkMode }) => {

  // Style for the valve
  const tankStyle = {
		position: 'absolute',
    top: `${coords.y}%`,
    left: `${coords.x}%`,
    height: `${dimensions.y}%`,
    width: `${dimensions.x}%`,
    border: '1px solid var(--main-fg-color)',
    borderRadius: '1rem'
  };

  return (
    <div style={tankStyle}></div>
  );
};

export default Tank;
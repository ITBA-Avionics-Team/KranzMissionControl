
const TankDynamicSize = ({ coords, dimensions }) => {
  // Style for the valve
  const style = {
    stroke: 'var(--main-fg-color)',
    fill: 'var(--main-fg-color)',
    strokeWidth: '1%',

		position: 'absolute',
    top: `${coords.y}%`,
    left: `${coords.x}%`,
    width: `${dimensions.x}%`,
    height: `${dimensions.y}%`,
    // transform: `scale(${dimensions.x}, ${dimensions.y})`
  };

  return (
  <svg style={style} xmlns="http://www.w3.org/2000/svg" viewBox={`0 0 ${dimensions.x} ${dimensions.y}`} version="1.1" >
    <rect fill="none" height={dimensions.y - 2*dimensions.y/10} rx={dimensions.y / 10} ry={dimensions.y / 10} width={dimensions.x - 2*dimensions.x/10} x="0" y="0" />
  </svg>
  );
};

export default TankDynamicSize;
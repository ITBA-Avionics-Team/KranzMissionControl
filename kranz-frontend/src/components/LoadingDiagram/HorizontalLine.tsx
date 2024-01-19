
const HorizontalLine = ({ coords, dimensions }) => {
  // Style for the valve
  const style = {
    stroke: 'var(--main-fg-color)',
    fill: 'var(--main-fg-color)',
    strokeWidth: '5%',

		position: 'absolute',
    top: `${coords.y}%`,
    left: `${coords.x}%`,
    width: `${dimensions.x}%`,
  };

  return (
  <svg style={style} xmlns="http://www.w3.org/2000/svg" viewBox="0 0 100 100" baseProfile="full" version="1.1" >
    <line x1="0" x2="100" y1="0" y2="0" />
  </svg>
  );
};

export default HorizontalLine;
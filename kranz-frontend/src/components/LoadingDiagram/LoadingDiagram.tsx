import tankSvg from '../../assets/loading-diagram/tank_image.svg'
import verticalLineSvg from '../../assets/loading-diagram/vertical_line.svg'
import horizontalLineSvg from '../../assets/loading-diagram/horizontal_line.svg'
import PressureSensor from './PressureSensor';
import TemperatureSensor from './TemperatureSensor';
import Umbrilical from './Umbrilical';
import Valve from './Valve';
import Image from './Image';
import Text from './Text';
import { useEffect, useRef, useState } from 'react';
import FloatingText from '../FloatingText';

const LoadingDiagram = ({ systemStatus }) => {

  // Style for the valve
  const containerStyle = {
		position: 'relative',
		width: '1999px',
		height: '1500px'
  };

	const [cursorText, setCursorText] = useState("CursorText");

	const elementRef = useRef(null);
	const [cursorPos, setCursorPos] = useState({ x: 0, y: 0 });
	const [cursorPosOffset, setOffset] = useState({ top: 0, left: 0 });
	useEffect(() => {
		if (elementRef.current) {
				const rect = elementRef.current.getBoundingClientRect();
				setOffset({
						top: rect.top + window.scrollY,
						left: rect.left + window.scrollX
				});
		}
	}, []);
	const handleMouseMove = (e) => {
		setCursorPos({ x: e.clientX - cursorPosOffset.left + 15, y: e.clientY - cursorPosOffset.top +15});
	};

	const openValveCommandDialog = (valveId) => {
		console.log("ValveOmmand:" + valveId);
  };

	const displayNameOnCursor = (componentId) => {
		setCursorText(componentId);
	}

	const clearCursorName = (componentId) => {
		setCursorText("");
	}


  return (
    <div ref={elementRef} style={containerStyle} onMouseMove={handleMouseMove} >
			<FloatingText cursorPos={cursorPos} text={cursorText}/>

			<Valve name="tank_depress_vent_valve" coords={{x:17, y:10}} status={systemStatus.on_board.tank_depress_vent_valve_open} onMouseMove={handleMouseMove} onClick={openValveCommandDialog} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Image src={verticalLineSvg} coords={{x:14, y:13}} dimensions={{x: 7, y: 2}}/>
			<TemperatureSensor name="tank_depress_vent_temp_celsius" coords={{x:17, y:5}} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Text value={systemStatus.on_board.tank_depress_vent_temp_celsius +  "°C"} coords={{x:18.5, y:5}}/>
			<Image src={tankSvg} coords={{x:15, y:15}} dimensions={{x: 5, y: 20}}/>
			<Text value={"NOX"} coords={{x:16.6, y:20}}/>
			<TemperatureSensor name="tank_temperature" coords={{x:20, y:20}} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Text value={systemStatus.on_board.tank_temp_celsius +  "°C"} coords={{x:21.5, y:20}}/>
			<PressureSensor name="tank_pressure" coords={{x:20, y:25}} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Text value={systemStatus.on_board.tank_pressure_psi +  " PSI"} coords={{x:21.5, y:25}}/>

			<Image src={verticalLineSvg} coords={{x:14, y:35}} dimensions={{x: 7, y: 8}}/>
			<Image src={horizontalLineSvg} coords={{x:17.5, y:35}} dimensions={{x: 10, y: 9}}/>
			<Image src={verticalLineSvg} coords={{x:14, y:46}} dimensions={{x: 7, y: 2}}/>

			<Valve name="engine_valve" coords={{x:17, y:43}} status={systemStatus.on_board.engine_valve_open} onClick={openValveCommandDialog} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Image src={tankSvg} coords={{x:15, y:48}} dimensions={{x: 5, y: 20}}/>
			<Text value={"Combustion"} coords={{x:15, y:53}}/>
			<TemperatureSensor name="combustion_chamber_temperature" coords={{x:20, y:53}} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Text value={systemStatus.on_board.combustion_chamber_temp_celsius + "°C"} coords={{x:21.5, y:53}}/>
			<PressureSensor name="combustion_chamber_pressure" coords={{x:20, y:59}} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Text value={systemStatus.on_board.combustion_chamber_pressure_psi + " PSI"} coords={{x:21.5, y:59}}/>

			<Umbrilical name="umbrilical" coords={{x:28, y:37}} status={systemStatus.launchpad.umbrilical_connected} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>

			<Image src={horizontalLineSvg} coords={{x:30.5, y:35}} dimensions={{x: 7, y: 9}}/>
			<Image src={horizontalLineSvg} coords={{x:40, y:35}} dimensions={{x: 4, y: 9}}/>
			<Image src={verticalLineSvg} coords={{x:33, y:39.4}} dimensions={{x: 1, y: 5}}/>
			<Valve name="loading_valve" coords={{x:38, y:38}} status={systemStatus.launchpad.loading_valve_open} onClick={openValveCommandDialog} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Valve name="loading_depress_vent_valve" coords={{x:33, y:45}} status={systemStatus.launchpad.loading_depress_vent_valve_opem} onClick={openValveCommandDialog} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
    </div>
  );
};

export default LoadingDiagram;
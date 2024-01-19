import noxTank from '../../assets/react.svg'
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
		setCursorPos({ x: e.clientX - cursorPosOffset.left, y: e.clientY - cursorPosOffset.top });
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

			<Valve name="tank_depress_vent_valve" coords={{x:15, y:10}} status={systemStatus.on_board.tank_depress_vent_valve_open} onMouseMove={handleMouseMove} onClick={openValveCommandDialog} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<TemperatureSensor name="tank_depress_vent_temp_celsius" coords={{x:15, y:5}} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Text value={systemStatus.on_board.tank_depress_vent_temp_celsius +  "°C"} coords={{x:15, y:0}}/>
			<Image src={noxTank} coords={{x:15, y:15}}/>
			<TemperatureSensor name="tank_temperature" coords={{x:20, y:20}} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Text value={systemStatus.on_board.tank_temp_celsius +  "°C"} coords={{x:25, y:20}}/>
			<PressureSensor name="tank_pressure" coords={{x:20, y:25}} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Text value={systemStatus.on_board.tank_pressure_psi +  " PSI"} coords={{x:25, y:25}}/>

			<Image src={noxTank} coords={{x:15, y:45}}/>
			<Valve name="engine_valve" coords={{x:15, y:35}} status={systemStatus.on_board.engine_valve_open} onClick={openValveCommandDialog} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<TemperatureSensor name="combustion_chamber_temperature" coords={{x:20, y:40}} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Text value={systemStatus.on_board.combustion_chamber_temp_celsius + "°C"} coords={{x:25, y:40}}/>
			<PressureSensor name="combustion_chamber_pressure" coords="{x:20, y:45}" onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Text value={systemStatus.on_board.combustion_chamber_pressure_psi + " PSI"} coords={{x:25, y:45}}/>

			<Umbrilical name="umbrilical" coords={{x:25, y:30}} status={systemStatus.launchpad.umbrilical_connected} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>

			<Valve name="loading_valve" coords={{x:30, y:30}} status={systemStatus.launchpad.loading_valve_open} onClick={openValveCommandDialog} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
			<Valve name="loading_depress_vent_valve" coords={{x:35, y:40}} status={systemStatus.launchpad.loading_depress_vent_valve_opem} onClick={openValveCommandDialog} onMouseEnter={displayNameOnCursor} onMouseLeave={clearCursorName}/>
    </div>
  );
};

export default LoadingDiagram;
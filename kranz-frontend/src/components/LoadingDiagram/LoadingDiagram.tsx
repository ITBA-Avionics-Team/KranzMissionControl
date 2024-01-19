import noxTank from '../../assets/react.svg'
import PressureSensor from './PressureSensor';
import TemperatureSensor from './TemperatureSensor';
import Umbrilical from './Umbrilical';
import Valve from './Valve';
import Image from './Image';
import Text from './Text';

const LoadingDiagram = ({ systemStatus }) => {

  // Style for the valve
  const containerStyle = {
		position: 'relative',
    // width: '50%', /* Relative width, 50% of the parent element's width */
	  // height: '75%', /* Relative height, 75% of the parent element's height */
		width: '1999px',
		height: '1500px'
  };

	// console.log(systemStatus);

	const openValveCommandDialog = (valveId) => {
		console.log("ValveOmmand:" + valveId);
  };

	const displayNameOnCursor = (componentId) => {
		console.log(componentId);
	
	}

  return (
    <div style={containerStyle}>
			<Valve name="tank_depress_vent_valve" coords={{x:15, y:10}} status={systemStatus.on_board.tank_depress_vent_valve_open} onClick={openValveCommandDialog} onHover={displayNameOnCursor}/>
			<TemperatureSensor name="tank_depress_vent_temp_celsius" coords={{x:15, y:5}} onHover={displayNameOnCursor}/>
			<Text value={systemStatus.on_board.tank_depress_vent_temp_celsius +  "°C"} coords={{x:15, y:0}}/>
			<Image src={noxTank} coords={{x:15, y:15}}/>
			<TemperatureSensor name="tank_temperature" coords={{x:20, y:20}} onHover={displayNameOnCursor}/>
			<Text value={systemStatus.on_board.tank_temp_celsius +  "°C"} coords={{x:25, y:20}}/>
			<PressureSensor name="tank_pressure" coords={{x:20, y:25}} onHover={displayNameOnCursor}/>
			<Text value={systemStatus.on_board.tank_pressure_psi +  " PSI"} coords={{x:25, y:25}}/>

			<Image src={noxTank} coords={{x:15, y:45}}/>
			<Valve name="engine_valve" coords={{x:15, y:35}} status={systemStatus.on_board.engine_valve_open} onClick={openValveCommandDialog} onHover={displayNameOnCursor}/>
			<TemperatureSensor name="combustion_chamber_temperature" coords={{x:20, y:40}} onHover={displayNameOnCursor}/>
			<Text value={systemStatus.on_board.combustion_chamber_temp_celsius + "°C"} coords={{x:25, y:40}}/>
			<PressureSensor name="combustion_chamber_pressure" coords="{x:20, y:45}" onHover={displayNameOnCursor}/>
			<Text value={systemStatus.on_board.combustion_chamber_pressure_psi + " PSI"} coords={{x:25, y:45}}/>

			<Umbrilical name="umbrilical" coords={{x:25, y:30}} status={systemStatus.launchpad.umbrilical_connected} onHover={displayNameOnCursor}/>

			<Valve name="loading_valve" coords={{x:30, y:30}} status={systemStatus.launchpad.loading_valve_open} onClick={openValveCommandDialog} onHover={displayNameOnCursor}/>
			<Valve name="loading_depress_vent_valve" coords={{x:35, y:40}} status={systemStatus.launchpad.loading_depress_vent_valve_opem} onClick={openValveCommandDialog} onHover={displayNameOnCursor}/>
    </div>
  );
};

export default LoadingDiagram;
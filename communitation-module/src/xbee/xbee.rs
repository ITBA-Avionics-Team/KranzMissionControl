use std::{time::Duration, io, sync::{mpsc, Arc, Mutex}};

use rocket::{serde::Deserialize, outcome::IntoOutcome};
use serde::Serialize;
use tokio::sync::{broadcast::{Sender, error::SendError}};

#[derive(Debug)]
#[derive(Clone)]
#[derive(Serialize)]
pub struct FlightComputersStatus {
    altium_ok: bool,
    ada_ok: bool
}

#[derive(Debug)]
#[derive(Clone)]
#[derive(Serialize)]
pub struct OnBoardSystemState {
    connection_status: String,
    tank_pressure_bar: f32,
    tank_temp_celsius: f32,
    tank_depress_vent_temp_celsius: f32,
    tank_depress_vent_valve_open: bool,
    engine_valve_open: bool,
    OBEC_battery_voltage: f32,
    flight_computers_status: FlightComputersStatus
}

#[derive(Debug)]
#[derive(Clone)]
#[derive(Serialize)]
pub struct LaunchpadSystemState {
    connection_status: String,
    load_line_pressure: f32,
    loading_valve_open: bool,
    loading_depress_vent_valve_opem: bool,
    umbrilical_connected: bool,
    igniter_continuity_ok: bool,
    external_vent_as_default: bool,
    LC_battery_voltage: f32,
}

#[derive(Debug)]
#[derive(Clone)]
#[derive(Serialize)]
pub struct WeatherData {
    wind_speed_knt: u8
}

#[derive(Debug)]
#[derive(Clone)]
#[derive(Serialize)]
pub struct SystemState {
    on_board: OnBoardSystemState,
    launchpad: LaunchpadSystemState,
    weather_data: WeatherData
}

impl SystemState {
    fn from_message(message: String) -> Result<SystemState, String> {
        let sensor_data_byte = message.chars().nth(28).ok_or("Failed to get sensor data byte from SystemMessage")? as u8;
        let obec_connection_ok = (sensor_data_byte as u8) & 0b01 == 1;
        let tank_depress_vent_valve_open = (sensor_data_byte as u8) & 0b10 == 1;
        let engine_valve_open = (sensor_data_byte as u8) & 0b100 == 1;
        let loading_valve_open = (sensor_data_byte as u8) & 0b1000 == 1;
        let loading_depress_vent_valve_open = (sensor_data_byte as u8) & 0b10000 == 1;
        let umbrilical_connected = (sensor_data_byte as u8) & 0b100000 == 1;
        let igniter_continuity_ok = (sensor_data_byte as u8) & 0b1000000 == 1;
        let external_vent_as_default = (sensor_data_byte as u8) & 0b1000000 == 1;

        
        Ok(SystemState {
            on_board: OnBoardSystemState {
                connection_status: String::from("ok"),
                tank_pressure_bar: message[4..8].parse::<f32>().map_err(|_| "Failed to parse tank pressure from SystemMessage")?,
                tank_temp_celsius: message[8..12].parse::<f32>().map_err(|_| "Failed to parse tank temperature from SystemMessage")?,
                tank_depress_vent_temp_celsius: message[12..16].parse::<f32>().map_err(|_| "Failed to parse tank depress vent temperature from SystemMessage")?,
                tank_depress_vent_valve_open: tank_depress_vent_valve_open,
                engine_valve_open: engine_valve_open,
                OBEC_battery_voltage: message[20..24].parse::<f32>().map_err(|_| "Failed to parse OBEC battery voltage from SystemMessage")?,
                flight_computers_status: FlightComputersStatus {
                    ada_ok: true,
                    altium_ok: true 
                }
            },
            launchpad: LaunchpadSystemState {
                connection_status: String::from("ok"), // TODO: Verify?
                load_line_pressure: message[16..20].parse::<f32>().map_err(|_| "Failed to parse load line pressure from SystemMessage")?,
                loading_valve_open: loading_valve_open,
                loading_depress_vent_valve_opem: loading_depress_vent_valve_open,
                umbrilical_connected: umbrilical_connected,
                igniter_continuity_ok: igniter_continuity_ok,
                external_vent_as_default: external_vent_as_default,
                LC_battery_voltage: message[24..28].parse::<f32>().map_err(|_| "Failed to parse LC battery voltage from SystemMessage")?
            },
            weather_data: WeatherData {
                wind_speed_knt: message[29..32].parse::<u8>().map_err(|_| "Failed to parse wind speed from SystemMessage")?
            }
        })
    }
}


pub fn send_command(raw_content: &str) {
  println!("Sending xbee message: {:?}", raw_content)
}


pub fn listen_for_messages(serial_port_name: &str, baud_rate: u32, data_sender: Arc<Sender<Arc<SystemState>>>) {
  // loop {}

    let port = serialport::new(serial_port_name, baud_rate)
        .timeout(Duration::from_millis(10))
        .open();

    match port {
        Ok(mut port) => {
            let mut serial_buf: Vec<u8> = vec![0; 1000];
            println!("Receiving data on {} at {} baud:", &serial_port_name, &baud_rate);
            loop {
                match port.read(serial_buf.as_mut_slice()) {
                    Ok(t) => {
                      let buffer_message = serial_buf[..t].into_iter().map(|&u|u as char).collect();
                      if (t > 0) {
                        let parsed_state_result = SystemState::from_message(buffer_message);
                        match parsed_state_result {
                            Ok(system_state) => {
                                let system_state: Arc<SystemState> = Arc::new(system_state);
                                let res = data_sender.send(system_state);
                                match res {
                                    Ok(sent_size) => {
                                        print!("Successful send");
                                    },
                                    Err(err) => {
                                        print!("{}", err.to_string());
                                    }
                                }
                            },
                            Err(error_msg) => {
                                println!("{}", error_msg);
                            }
                        }
                        
                      }
                      
                    },
                    Err(ref e) if e.kind() == io::ErrorKind::TimedOut => (),
                    Err(e) => eprintln!("{:?}", e),
                }
            }
        }
        Err(e) => {
            println!("Failed to open \"{}\". Error: {}", serial_port_name, e);
            // ::std::process::exit(1);
        }
    }
}

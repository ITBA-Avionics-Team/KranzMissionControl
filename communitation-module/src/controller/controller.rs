use std::{borrow::Borrow, sync::Arc};

use rocket_ws::{WebSocket, Channel, Message};
use rocket::{serde::{Deserialize, json::Json}, futures::SinkExt, response::status};

use crate::{AppState, xbee};

#[get("/system_status")]
pub fn system_status(state: &rocket::State<AppState>, ws: WebSocket) -> Channel<'_> {
    println!("Opened new websocket connection, creating tokio thread for xbee messages");
    let mut status_broadcast_clone = state.system_status_broadcast.subscribe();
    ws.channel(move |mut stream| Box::pin(async move {
        println!("Opened ws channel, spawning tokio thread");
        loop {
            let message = {
                // let mut guard = status_broadcast_clone.lock().unwrap();
                status_broadcast_clone.recv().await
            };

            if let Ok(message) = message {
                let system_state = message.as_ref();
                let _ = stream.send(Message::from(serde_json::to_string(system_state).unwrap())).await;
            }
        }
        Ok(())
    }))
}

#[post("/command/<destination>", format = "application/json", data = "<command>")]
pub fn command_lc(destination: &str, command: Json<Command<'_>>) -> String {
    crate::xbee::xbee::send_command(command.raw_content);
    String::from("Command received and sent to destionation ") + destination + &String::from("\n")
}

#[derive(Deserialize)]
#[derive(Debug)]
#[serde(crate = "rocket::serde")]
struct Command<'r> {
    raw_content: &'r str,
}

// struct SystemState {
//     tank_pressure_psi: f64,
//     tank_temp_celsius: f64,
//     tank_depress_vent_temp_celsius: f64,
//     tank_depress_vent_valve_open: bool,
//     engine_valve_open: bool,
//     OBEC_battery_voltage: f32,
//     igniter_continuity: bool,
//     flight_computers_ok: {
//         altium: bool,
//         ada: bool
//     }
//     load_line_pressure: f64,
//     loading_valve_open: bool,
//     loading_depress_vent_valve_opem: bool,
//     umbrilical_connected: bool,
//     LC_battery_voltage: f32,
//     weather_data: {
//         wind_speed_knt: u16,
//         temperature_celsius: f32
//     }
// }
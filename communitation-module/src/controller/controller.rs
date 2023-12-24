use rocket_ws::{WebSocket, Stream};
use rocket::serde::{Deserialize, json::Json};

#[get("/system_status")]
pub fn system_status(ws: WebSocket) -> Stream!['static] {
    ws.stream(|io| io)
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

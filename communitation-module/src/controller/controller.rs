use rocket_ws::{WebSocket, Stream};
use rocket::serde::{Deserialize, json::Json};

#[get("/system_status")]
pub fn system_status(ws: WebSocket) -> Stream!['static] {
    ws.stream(|io| io)
}

// Note: without the `..` in `opt..`, we'd need to pass `opt.emoji`, `opt.name`.
//
// Try visiting:
//   http://127.0.0.1:8000/?emoji
//   http://127.0.0.1:8000/?name=Rocketeer
//   http://127.0.0.1:8000/?lang=ру
//   http://127.0.0.1:8000/?lang=ру&emoji
//   http://127.0.0.1:8000/?emoji&lang=en
//   http://127.0.0.1:8000/?name=Rocketeer&lang=en
//   http://127.0.0.1:8000/?emoji&name=Rocketeer
//   http://127.0.0.1:8000/?name=Rocketeer&lang=en&emoji
//   http://127.0.0.1:8000/?lang=ru&emoji&name=Rocketeer
#[post("/command", format = "application/json", data = "<command>")]
pub fn command(command: Json<Command<'_>>) -> String {
    let mut greeting = String::new();

    println!("{:?}", command);
    "Command received\n".to_string()
}

#[derive(Deserialize)]
#[derive(Debug)]
#[serde(crate = "rocket::serde")]
struct Command<'r> {
    description: &'r str,
    complete: bool
}

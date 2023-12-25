#[macro_use] extern crate rocket;
mod controller;
mod xbee;

use std::sync::mpsc;

use xbee::xbee::{listen_for_messages, LCMessage};

struct AppState<'a> {
    xbee_sender: mpsc::Sender<LCMessage<'a>>,
}

#[launch]
fn rocket() -> _ {
    let (xbee_sender, xbee_receiver) = mpsc::channel();
    std::thread::spawn(|| listen_for_messages("/dev/ttys004", 0, xbee_receiver));

    rocket::build()
        .manage(AppState { xbee_sender })
        .mount("/", routes![controller::controller::command_lc, controller::controller::system_status])
}

#[macro_use] extern crate rocket;
mod controller;
mod xbee;

use std::sync::{mpsc, Arc};

use rocket::futures::lock::Mutex;
use xbee::xbee::{listen_for_messages, LCMessage};

struct AppState {
    xbee_data_receiver: Arc<Mutex<mpsc::Receiver<LCMessage>>>,
}

#[launch]
fn rocket() -> _ {
    let (xbee_data_tx, xbee_data_rx) = mpsc::channel();
    let xbee_data_receiver = Arc::new(Mutex::new(xbee_data_rx));
    
    std::thread::spawn(|| listen_for_messages("/dev/ttys004", 0, xbee_data_tx));

    rocket::build()
        .manage(AppState { xbee_data_receiver })
        .mount("/", routes![controller::controller::command_lc, controller::controller::system_status])
}

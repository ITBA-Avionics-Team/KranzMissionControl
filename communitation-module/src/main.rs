#[macro_use]
extern crate rocket;
mod controller;
mod xbee;

use std::{
    borrow::{Borrow, BorrowMut},
    sync::{Arc, Mutex},
};
use tokio::sync::{broadcast::{self, Receiver, Sender}};

use rocket::{Build, Rocket};
use xbee::xbee::{listen_for_messages, SystemState};

struct AppState {
    system_status_broadcast: Arc<Sender<Arc<SystemState>>>,
}

#[launch]
#[tokio::main]
async fn rocket() -> _ {
    let (sender, _) = broadcast::channel::<Arc<SystemState>>(10);
    // let system_status_broadcast = sender.subscribe();
    let system_status_broadcast_sender = Arc::new(sender);
    let sender_clone = Arc::clone(&system_status_broadcast_sender);
    
    tokio::spawn(async move { listen_for_messages("/dev/ttys008", 0, Arc::clone(&system_status_broadcast_sender)) });
    rocket::build()
        .manage(AppState {
            system_status_broadcast: sender_clone,
        })
        .mount(
            "/",
            routes![
                controller::controller::command_lc,
                controller::controller::system_status
            ],
        )
}

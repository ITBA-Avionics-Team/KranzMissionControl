#[macro_use] extern crate rocket;
mod controller;
mod xbee;

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![controller::controller::command_lc, controller::controller::system_status])
}
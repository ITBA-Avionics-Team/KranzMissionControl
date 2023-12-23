#[macro_use] extern crate rocket;
mod controller;

#[launch]
fn rocket() -> _ {
    rocket::build()
        .mount("/", routes![controller::controller::command, controller::controller::system_status])
}
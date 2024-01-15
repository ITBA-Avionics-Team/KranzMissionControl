use std::{time::{Duration, self}, thread, env, process::exit};


const OBEC_STATUS_MESSAGE: &[u8] = b"001414.125.24.12)|";

const SYSTEM_STATUS_MESSAGE: &[u8] = b"STBY001410.125.200144.124.20?008|";
const SYSTEM_STATUS_MESSAGE_ABORT: &[u8] = b"ABRT001410.125.200144.124.20@008|";

fn main() {
    let write_port_name: &str;
    let args: Vec<String> = env::args().collect();
    if args.len() > 1 {
        // Get the first argument (excluding the program name)
        write_port_name = &args[1];
    } else {
        println!("No argument provided to specify port. Example: /dev/ttys007");
        exit(1)
    }

    let baud_rate = 0;

    // Open the write port
    let mut write_port = serialport::new(write_port_name, baud_rate)
        .timeout(Duration::from_millis(10))
        .open()
        .expect("Failed to open write port");

    // // Open the read port
    // let read_port = serialport::new(read_port_name, baud_rate)
    //     .timeout(Duration::from_millis(10))
    //     .open()
    //     .expect("Failed to open read port");


    let mut abort = false;
    loop {
        if (!abort) {
            write_port.write_all(&SYSTEM_STATUS_MESSAGE);
        }
        else {
            write_port.write_all(&SYSTEM_STATUS_MESSAGE_ABORT);
        }
        abort = !abort;
        let one_second = time::Duration::from_millis(1000);
        thread::sleep(one_second);
    }

    // Close the ports when done
    // write_port.close().expect("Failed to close write port");
    // read_port.close().expect("Failed to close read port");
}
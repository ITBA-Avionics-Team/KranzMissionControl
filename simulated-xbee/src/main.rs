use std::{time::{Duration, self}, thread};

fn main() {
    let write_port_name = "/dev/ttys003"; // Replace with the correct port
    // let read_port_name = "/dev/ttys01";  // Replace with the correct port
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


    let buff: [u8; 28] = [b'0', b'0', b'1', b'4', b'1', b'0', b'.', b'1', b'2', b'5', b'.', b'2', b'0', b'0', b'1', b'4', b'4', b'.', b'1', b'2', b'4', b'.', b'2', b'0', b'?', b'0', b'0', b'8'];
    loop {
        write_port.write_all(&buff);
        let one_second = time::Duration::from_millis(1000);
        thread::sleep(one_second);
    }

    // Close the ports when done
    // write_port.close().expect("Failed to close write port");
    // read_port.close().expect("Failed to close read port");
}
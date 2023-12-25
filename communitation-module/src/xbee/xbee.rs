use std::{time::Duration, io, sync::mpsc};

#[derive(Debug)]
pub struct LCMessage<'r> {
    raw_content: &'r str,
}



pub fn send_command(raw_content: &str) {
  println!("Sending xbee message: {:?}", raw_content)
}

pub fn listen_for_messages(serial_port_name: &str, baud_rate: u32, rx: mpsc::Receiver<LCMessage>) {
  // loop {}

    let port = serialport::new(serial_port_name, baud_rate)
        .timeout(Duration::from_millis(10))
        .open();

    match port {
        Ok(mut port) => {
            let mut serial_buf: Vec<u8> = vec![0; 1000];
            println!("Receiving data on {} at {} baud:", &serial_port_name, &baud_rate);
            loop {
                match port.read(serial_buf.as_mut_slice()) {
                    Ok(t) => {
                      let bufferMessage: String = serial_buf[..t].into_iter().map(|&u|u as char).collect();
                      println!("{:?}", bufferMessage)
                    },
                    Err(ref e) if e.kind() == io::ErrorKind::TimedOut => (),
                    Err(e) => eprintln!("{:?}", e),
                }
            }
        }
        Err(e) => {
            println!("Failed to open \"{}\". Error: {}", serial_port_name, e);
            // ::std::process::exit(1);
        }
    }
}

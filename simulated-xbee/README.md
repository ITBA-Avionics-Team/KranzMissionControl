# Simulated XBee device
This program simply writes to a serial port assuming that it has been previously configured as the write end of a virtual serial port tunnel.
You can create such a virtual serial port using the `socat` utility by running the command `socat -d -d pty,raw,echo=0 pty,raw,echo=0`

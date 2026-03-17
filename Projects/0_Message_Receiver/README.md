# [0] TCP/UDP Message Receiver

TCP and UDP are Transport Layer protocols that are core components of the Internet protocol suite.

UDP is a simple protocol. It provides guaranteed message integrity but it does not guarantee that the message will be delivered due to it being a stateless protocol. UDP is generally used for real-time systems where waiting for packages to arrive is not an option. Certain Application Layer protocols like DNS are built on top of UDP.

TCP, on the other hand, is reliable, oredered, and performs error checking as well. Being connection-oriented, it first establishes a 3-way handshake procedure to connect to the destination before sending a data stream to the listener. Among the Application Layer protocols built on top of TCP are TLS, SMTP, etc.

## UDP Receiver
<p align="center">
    <img src="../../images/UDP Receiver Diagram.png">
</p>

The logic behind an UDP receiver is simple: create a socket and bind it to an IP address and a port, in this case _localhost_, and port 6913.

Once bound, open a listener inside a loop waiting to receive data. This listener will have a buffer size of at least 1024 bytes, enough room to receive a few words sent from another terminal session.

To send a message run the [send] function in the receiver file, to receive, use [receive]. Any message typed in the terminal during send will arrive to a receiver so long as the receiver is up and running, listening for data streams.

## TCP Receiver
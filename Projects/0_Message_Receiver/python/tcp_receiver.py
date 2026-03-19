import socket

def receiver():
    host = ''
    port = 1369
    # Open socket in TCP mode and bind it to localhost:1369
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sck:
        sck.bind((host, port))
        # Listen and accept connections
        sck.listen(1)
        conn, address = sck.accept()

        data = None
        print('Listening for messages...')
        while data != b'exit':
            # Print any data received
            data = conn.recv(1024)
            print('>', data.decode())


def sender():
    localhost = '127.0.0.1'
    port = 1369
    # Open socket in TCP and connect to localhost:1369
    with socket.socket(socket.AF_INET, socket.SOCK_STREAM) as sck:
        sck.connect((localhost, port))
        print('Connection successful')

        message = None
        while message != 'exit':
            # Send input data
            message = input('Insert message to send (send `exit` to stop): ')
            sck.sendall(message.encode('utf-8'))


if __name__ == '__main__':
    option = input('[send] or [receive] data? ')
    match option:
        case 'send':
            sender()
        case 'receive':
            receiver()
        case _:
            exit()
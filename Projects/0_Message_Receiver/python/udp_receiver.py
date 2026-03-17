import socket

def receiver():
    localhost = '127.0.0.1'
    port = 6913
    sck = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    sck.bind((localhost, port))

    data = None
    print('Listening for messages')
    while data != b'exit':
        data, _address = sck.recvfrom(1024)
        print(data.decode())

def sender():
    localhost = '127.0.0.1'
    port = 6913
    sck = socket.socket(socket.AF_INET, socket.SOCK_DGRAM)
    
    message = None
    while message != 'exit':
        message = input('Insert message to send (send `exit` to stop): ')
        sck.sendto(message.encode('utf-8'), (localhost, port))



if __name__ == '__main__':
    option = input('[send] or [receive] data? ')
    match option:
        case 'send':
            sender()
        case 'receive':
            receiver()
        case _:
            exit()


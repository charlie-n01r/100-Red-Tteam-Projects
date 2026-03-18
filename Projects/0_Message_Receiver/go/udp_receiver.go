package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func receiver() error {
	listener, err := net.ListenPacket("udp4", ":6913")
	if err != nil {
		return err
	}
	defer listener.Close()

	dataBuff := make([]byte, 1024)
	for {
		n, _, err := listener.ReadFrom(dataBuff)
		if err != nil {
			return err
		}
		message := strings.TrimSpace(string(dataBuff[:n]))
		if message == "exit" {
			break
		}

		fmt.Println(message)
	}

	return nil
}

func sender() error {
	socket, err := net.Dial("udp4", "127.0.0.1:6913")
	if err != nil {
		return err
	}
	defer socket.Close()

	var message []byte

	for {
		fmt.Scan(&message)
		_, err = socket.Write(message)
		if err != nil {
			return err
		}

		if string(message) == "exit" {
			break
		}
	}

	return nil
}

func main() {
	var option string
	fmt.Println("[send] or [receive] data?")
	fmt.Scan(&option)
	switch option {
	case "send":
		err := sender()
		if err != nil {
			fmt.Println(err)
		}
	case "receive":
		err := receiver()
		if err != nil {
			fmt.Println(err)
		}
	default:
		os.Exit(0)
	}
}

package main

import (
	"fmt"
	"net"
	"os"
	"strings"
)

func receiver() error {
	// Create a TCP listener on localhost:1369
	listener, err := net.Listen("tcp", ":1369")
	if err != nil {
		return err
	}
	// Accept incoming connection
	conn, err := listener.Accept()
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("Listening for messages...")
	for {
		dataBuff := make([]byte, 1024)
		// Read any incoming messages and print them
		_, err := conn.Read(dataBuff)
		if err != nil {
			return err
		}

		message := strings.TrimSpace(string(dataBuff))
		fmt.Println(message)
		if message == "exit" {
			break
		}
	}

	return nil
}

func sender() error {
	// Create a connection and dial localhost:1369
	conn, err := net.Dial("tcp", "127.0.0.1:1369")
	if err != nil {
		return err
	}
	defer conn.Close()

	fmt.Println("Connection successful")
	var message []byte
	for {
		fmt.Print("Insert message to send (send `exit` to stop): ")
		fmt.Scan(&message)
		// Write input message to listener
		_, err := conn.Write(message)
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

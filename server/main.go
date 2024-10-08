package main

import (
	"fmt"
	"net"
)

func handleConnection(client net.Conn) {
	defer client.Close()

	// copied example:
	// Example of reading and writing to the connection
	buffer := make([]byte, 1024)
	n, err := client.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	// Echo the data back to the client
	_, err = client.Write(buffer[:n])
	if err != nil {
		fmt.Println("Error writing:", err)
		return
	}
}
func createSocket() (net.Listener, error) {
	var socket net.Listener
	var err error

	socket, err = net.Listen("tcp", ":4000")
	if err != nil {
		fmt.Println("ERR: net.listen - ", err)
		return nil, err
	}

	fmt.Println("Socket listening on port 4000")
	return socket, nil
}

func acceptConnection(socket net.Listener) {
	var client net.Conn
	var err error

	for {
		client, err = socket.Accept()
		if err != nil {
			fmt.Println("ERR: could not connect to client -", err)
			continue
		}

		go handleConnection(client)

	}

}

func main() {
	var socket net.Listener
	var err error

	socket, err = createSocket()

	if err != nil {
		return
	}

	acceptConnection(socket)

}

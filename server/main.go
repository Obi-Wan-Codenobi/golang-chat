package main

import (
	"fmt"
	"net"
	"time"
)

type numberOfConnections struct {
	count       int
	updateCount chan int
}

func (c *numberOfConnections) countConnections() {
	for {
		// thread to update the count from the channel/pipe
		newCount := <-c.updateCount
		c.count = newCount
	}
}

func handleConnection(client net.Conn, c *numberOfConnections) {
	defer client.Close()
	c.updateCount <- c.count + 1

	buffer := make([]byte, 1024)
	n, err := client.Read(buffer)
	if err != nil {
		fmt.Println("Error reading:", err)
		return
	}

	time.Sleep(15 * time.Second)

	// Echo the data back to the client
	_, err = client.Write(buffer[:n])
	c.updateCount <- c.count - 1
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

func acceptConnection(socket net.Listener, c *numberOfConnections) {
	var client net.Conn
	var err error

	for {
		client, err = socket.Accept()
		if err != nil {
			fmt.Println("ERR: could not connect to client -", err)
			continue
		}

		go handleConnection(client, c)

	}

}

func main() {
	var socket net.Listener
	var err error

	socket, err = createSocket()

	if err != nil {
		return
	}

	// Makes channels
	c := &numberOfConnections{
		// This is essentially making the pipe for the threads to talk
		updateCount: make(chan int),
	}
	//checking count channel and updating
	go c.countConnections()

	go acceptConnection(socket, c)

	for {
		time.Sleep(5 * time.Second)
		//This may not be the best method to keep count since
		//	multiple threads update this value. Using a mutex
		//	would be better but I wanted to use routines to 
		//	get familar
		fmt.Println("Number of Active Connections:", c.count)
	}

}

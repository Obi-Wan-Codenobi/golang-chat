package main

import (
	"fmt"
	"net"
)

func handleConnection(server net.Conn){
    defer server.Close();
    
    // copied example:
    message := []byte("Hi there")
    // Echo the data back to the client
    _, err := server.Write(message);
    if err != nil {
        fmt.Println("Error writing:", err)
        return
    }

    // Example of reading and writing to the connection
    buffer := make([]byte, 1024)
    n, err := server.Read(buffer)
    if err != nil {
        fmt.Println("Error reading:", err)
        return
    }
    fmt.Println("Server response: ", string(buffer[:n]));

}
func createSocket() (net.Conn, error) {
    var server net.Conn; 
    var err error;

    server, err = net.Dial("tcp", "localhost:4000");
    if err != nil {
        fmt.Println("ERR: net.listen - ", err);
        return nil, err;
    }

    fmt.Println("Socket listening on port 4000");
    return server, nil;
} 


func main() {
    var serverConnection net.Conn;
    var err error; 

    serverConnection, err = createSocket();

    if err != nil{
        return;
    }

    handleConnection(serverConnection);


}

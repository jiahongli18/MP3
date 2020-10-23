package main

import (
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"strings"

	types "../structs"
	utils "../utils"
)

func main() {
	fmt.Println(utils.FetchFaultyNodes())
	//Declare a map to be used for storing the client's username and the associating TCP channel.
	var m = make(map[string]net.Conn)

	//Channel used to communicate between EXIT goroutine
	channel := make(chan string)
	go startServer(m)
	go exit(channel)
	signal := <-channel

	//upon receiving signal to exit, terminates process and channels with other clients as well by calling exitAllClients
	if signal == "EXIT" {
		exitAllClients(m)
		return
	}

}

//This function creates the TCP connection and listen on provided port for requests.
//It stores the username of the client's in a map as a key, and the channel as the associating value.
func startServer(m map[string]net.Conn) {
	username := ""

	PORT := ":" + utils.ReadStringFromConfig("port")
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Listening on port" + PORT + ". Please type 'EXIT' to quit.")
	}
	defer l.Close()

	// numOfNodes := utils.ReadIntFromConfig("numOfNodes")
	numOfFaultyNodes := utils.ReadIntFromConfig("numOfFaultyNodes")

	fmt.Println("There are this many faulty nodes: ")
	fmt.Println(numOfFaultyNodes)

	for {
		//The server accepts and begins to interact with TCP client
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		m[username] = c

		//Call a goroutine to handle the communication so that the server can support handling multiple concurrent clients.
		go handleConnection(c, m)
	}
}

//Listens on incoming requests from each client, redirects message to the receiver by iterating through map to find the associating channel.
func handleConnection(c net.Conn, m map[string]net.Conn) {
	for {
		decoder := gob.NewDecoder(c) //initialize gob decoder
		message := new(types.Message)
		_ = decoder.Decode(message)

		// fmt.Println(message)
	}
}

//Send termination signal to all clients in the map to terminate
func exitAllClients(m map[string]net.Conn) {
	for _, receiverChannel := range m {
		encoder := gob.NewEncoder(receiverChannel)
		msg := types.Message{"0", 0, false}
		encoder.Encode(msg)
	}
}

//Read user input for "EXIT" command, upon receiving, send signal back to main thread to exit
func exit(channel chan string) {
	for {
		reader := bufio.NewReader(os.Stdin)
		var cmd string
		cmd, _ = reader.ReadString('\n')
		if strings.TrimSpace(cmd) == "EXIT" {
			fmt.Println("Server is exiting...")
			//Sends the termination signal to the main thread
			channel <- "EXIT"
			return
		}
	}
}

package main

import (
	"bufio"
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"net"
	"os"
	"reflect"
	"strings"

	types "./utils"
)

func main() {
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
	//Scan and Parse in line argument for the port number
	arguments := os.Args
	if len(arguments) == 1 {
		fmt.Println("Please provide port number")
		return
	}

	PORT := ":" + arguments[1]
	l, err := net.Listen("tcp", PORT)
	if err != nil {
		fmt.Println(err)
		return
	} else {
		fmt.Println("Listening on port" + PORT + ". Please type 'EXIT' to quit.")
	}
	defer l.Close()

	numOfConnectedNodes := 0
	numOfNodes := readConfig("numOfNodes")
	numOfFaultyNodes := readConfig("numOfFaultyNodes")
	roundNum := 1

	fmt.Println("There are this many faulty nodes: ")
	fmt.Println(numOfFaultyNodes)

	for {
		//The server accepts and begins to interact with TCP client
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		// netData, _ := bufio.NewReader(c).ReadString('\n')
		// username = netData
		m[username] = c

		fmt.Println("Client connected")
		numOfConnectedNodes += 1

		if numOfConnectedNodes == numOfNodes {
			fmt.Println("everyone is connected")
			setFaultyNodes(c, numOfFaultyNodes, numOfNodes, roundNum, m)
		}

		//Call a goroutine to handle the communication so that the server can support handling multiple concurrent clients.
		go handleConnection(c, m)
	}
}

func MapRandomKeyGet(mapI interface{}) interface{} {
	keys := reflect.ValueOf(mapI).MapKeys()
	fmt.Print(keys)
	return keys[rand.Intn(len(keys))].Interface()
}

//use randomization to find which nodes fail ( need numOfFaultyNodes amount ex. if numOfFaultyNodes is 2, then let's say nodes 1 and 2 fail)
func setFaultyNodes(c net.Conn, numOfFaultyNodes int, numOfNodes int, roundNum int, m map[string]net.Conn) {
	var faultyNodes []string

	for i := 0; i < numOfFaultyNodes; i++ {
		faultyNodes = append(faultyNodes, MapRandomKeyGet(m).(string))
	}

	fmt.Println(faultyNodes)

	for key, receiverChannel := range m {
		isFaulty := false
		for i := 0; i < len(faultyNodes); i++ {
			if key == faultyNodes[i] {
				fmt.Print(key)
				isFaulty = true
			}
		}
		encoder := gob.NewEncoder(receiverChannel)
		msg := types.Message{0, 0, isFaulty, true}
		encoder.Encode(msg)
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
		msg := types.Message{0, 0, false, true}
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

func readConfig(desiredInfo string) (data int) {
	// Open  jsonFile
	jsonFile, err := os.Open("./config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	config := types.Config{}
	json.Unmarshal(byteValue, &config)

	//return desired info
	if desiredInfo == "numOfNodes" {
		return len(config.Nodes)
	} else if desiredInfo == "numOfFaultyNodes" {
		return config.NumOfFaultyNodes
	} else {
		return 0
	}

}

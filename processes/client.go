package main

import (
	"encoding/gob"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net"
	"os"

	types "../utils"
)

func main() {
	arguments := os.Args
	address := readConfig("host:port")
	name := arguments[0]

	//connect to provided host:post via the net library
	CONNECT := address
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	username := name
	fmt.Fprintf(c, username+"\n")

	channel := make(chan string)

	go listen(c, channel)

	//Wait for input from channel or user input, if "EXIT" command is received, then terminate.
	//If not signal is received, then get user input.
	for {
		select {
		case signal := <-channel:
			if signal == "EXIT" {
				return
			}
		default:
			msg := types.Message{1, 0.45, false, false}
			encoder := gob.NewEncoder(c)
			_ = encoder.Encode(msg)
		}
	}
}

//The listen function waits on messages from the TCP server.
func listen(c net.Conn, channel chan string) {
	for {
		decoder := gob.NewDecoder(c) //initialize gob decoder
		//Decode message struct and print it
		message := new(types.Message)
		_ = decoder.Decode(message)

		if (*message == types.Message{0, 0, false, true}) {
			//If the server terminates, the functions sends the exit signal through the channel to the main thread
			c.Close()
			os.Exit(0)
			channel <- "EXIT"
		}
	}
}

func readConfig(desiredInfo string) (address string) {
	// Open  jsonFile
	jsonFile, err := os.Open("../config.json")
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
	if desiredInfo == "host:port" {
		return config.Host + ":" + config.DefaultPort
	} else {
		return ""
	}

	// for i := 0; i < len(config.Nodes); i++ {
	// 	fmt.Println(config.Nodes[i].Value)
	// }
}

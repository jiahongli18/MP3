package main

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"sync"

	client "./client"
	server "./server"
	types "./structs"
	utils "./utils"
)

func main() {
	arguments := os.Args
	address := utils.ReadStringFromConfig("host:port")
	username := arguments[1]

	go server.StartServer(arguments[1])
	initialize(arguments[1])

	//connect to provided host:post via the net library
	CONNECT := address
	c, err := net.Dial("tcp", CONNECT)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Fprintf(c, username+"\n")
	channel := make(chan string)
	go client.Listen(c, channel)

	client.Unicast(arguments[1])

	//Wait for input from channel or user input, if "EXIT" command is received, then terminate.
	//If not signal is received, then get user input.
	for {
		select {
		case signal := <-channel:
			if signal == "EXIT" {
				return
			}
		default:
			msg := types.Message{"0", 0.0, false}
			encoder := gob.NewEncoder(c)
			_ = encoder.Encode(msg)
		}
	}
}

//The function initializes the dialers of the TCP connections and uses a delay in order to avoid excessive dialing
func initialize(processNum string) {
	processIP, host := utils.FetchHostPort(processNum)
	ports := utils.FetchPorts()

	//loop through every port in the config.txt and create a TCP connection between current process' port and others
	for port := range ports {

		//keeps dialing until a successful connection was made
		for {
			// status := dial(host, ports[port], processIP)
			status := dial(host, ports[port], processIP)
			if status == "success" {
				break
			}

			fmt.Println("Awaiting connections...Retrying in 2 secs.")

			//create a delay a goroutine and waitgroups
			wg := new(sync.WaitGroup)
			go utils.Delay(2000, 2001, wg)

			wg.Add(1)
			wg.Wait()
		}
	}
}

//dials to every other processes
func dial(processPort string, port string, ip string) (status string) {
	if port != processPort {
		address := ip + ":" + port

		_, err := net.Dial("tcp", address)
		if err != nil {
			return "error"
		}
	}
	return "success"
}

package main

import (
	"fmt"
	"net"
	"os"
	"sync"

	"./Utils"
	"./processes"
)

//The function initializes the dialers of the TCP connections and uses a delay in order to avoid excessive dialing
func initialize(NodeNum string) {
	processIP, host, _ := Utils.FetchHostPort(NodeNum)
	ports := Utils.FetchPorts()

	//loop through every port in the config.txt and create a TCP connection between current process' port and others
	for port := range ports {
		//keeps dialing until a successful connection was made
		for {
			status := dial(host, ports[port], processIP)
			if status == "success" {
				break
			}
			fmt.Println("Awaiting connections...Retrying in 2 secs.")

			//create a delay a goroutine and waitgroups
			wg := new(sync.WaitGroup)
			go Utils.Delay(2000, 2001, wg)
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

func main() {
	arguments := os.Args
	NodeNum := arguments[1]
	_, _, initialState := Utils.FetchHostPort(NodeNum)
	//initmsg := Utils.Message{}

	//initmsg.State = initialState
 	//initmsg.R = 1

	//Starts the server for each process in a goroutine so that it can listen to the dialing
	go processes.StartServer(NodeNum)

	initialize(NodeNum)
	processes.Unicast_send(initialState, 2)

}

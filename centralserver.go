package main

import (
	"./utils"
	"bufio"
	"encoding/gob"
	"fmt"
	"net"
)

/* func initializeCentral() {
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
} */

//Central server is an independent server that connects to all the nodes
func main() {
	var m = make(map[string]net.Conn)
	_,_, numNodes, numFailures := Utils.FetchConfig()
	fmt.Print(numNodes, numFailures)
	go startServer(m)
	

}

func startServer(m map[string]net.Conn) {
	var averageQueue []float64
	p := &averageQueue
	
	//get port number from user input and listen in on that port for requests
	PORT := ":" + "8081"
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}else {
		fmt.Println("Listening on port" + PORT)
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}
		netData, _ := bufio.NewReader(c).ReadString('\n')
		username := netData
		m[username] = c
		fmt.Print(&c)

		//goroutine for handling requests made to server
		go handleConnection(c, p)
	}
}

//Listens on incoming requests from each client, redirects message to the receiver by iterating through map to find the associating channel.
//func handleConnection(c net.Conn, ports []string, stateQueue *[]float64) {
func handleConnection(c net.Conn, p *[]float64) {
		for {
			decoder := gob.NewDecoder(c) //initialize gob decoder
			//msg := new(Utils.Message)
			state := new(float64)
	
			_ = decoder.Decode(state)
			//unicast_receive(*state, stateQueue)
		}
		//c.Close()
	}

/* 	func unicast_receive(state float64, stateQueue []float64) {
		stateQueue = append(stateQueue, state)
		if state > 0 {
			fmt.Printf("\nReceived %f, system time is %s\n", state, time.Now().Format("Jan _2 15:04:05.000"))
			fmt.Print(stateQueue)
		}
	} */
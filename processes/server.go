package processes

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"

	"../Utils"
)

func unicast_receive(state float64, stateQueue []float64) {
	stateQueue = append(stateQueue, state)
	if state > 0 {
		fmt.Printf("\nReceived %f, system time is %s\n", state, time.Now().Format("Jan _2 15:04:05.000"))
		fmt.Print(stateQueue)
	}
}

func handleConnection(c net.Conn, stateQueue []float64) {
	//var stateQueue [10000]float64
	for {
		decoder := gob.NewDecoder(c) //initialize gob decoder
		state := new(float64)
		//message := new(Utils.Message)
		_ = decoder.Decode(state)
		unicast_receive(*state, stateQueue)

	}
	c.Close()
}

func StartServer(NodeNum string) {
	var stateQueue []float64
	_, port, _ := Utils.FetchHostPort(NodeNum)

	//get port number from user input and listen in on that port for requests
	PORT := ":" + port
	l, err := net.Listen("tcp4", PORT)
	if err != nil {
		fmt.Println(err)
		return
	}
	defer l.Close()

	for {
		c, err := l.Accept()
		if err != nil {
			fmt.Println(err)
			return
		}

		//goroutine for handling requests made to server
		go handleConnection(c, stateQueue)
	}
}

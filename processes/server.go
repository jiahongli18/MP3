package processes

import (
	"../Utils"
	"fmt"
	"net"
	"encoding/gob"
	"time"
)

//func unicast_receive(source string, state float64) {
//	fmt.Printf("\nReceived %f from node %s, system time is %s\nPlease enter a command: ", state, source, time.Now().Format("Jan _2 15:04:05.000"))
//}

func unicast_receive(state float64) {
	fmt.Printf("\nReceived %f, system time is %s\n", state, time.Now().Format("Jan _2 15:04:05.000"))
	}



func handleConnection(c net.Conn) {
	for {
		decoder := gob.NewDecoder(c) //initialize gob decoder
		state := new(float64)
		//message := new(Utils.Message)
		_ = decoder.Decode(state)
		unicast_receive(*state)
	}
	c.Close()
}

func StartServer(NodeNum string) {
	_,port := Utils.FetchHostPort(NodeNum)

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
		go handleConnection(c)
	}
}

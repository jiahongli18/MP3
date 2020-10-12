package processes

import (
	"encoding/gob"
	"fmt"
	"net"

	"../Utils"
)

//func unicast_receive(msg Utils.Message, pmsg *[]Utils.Message, stateQueue *[]float64) {
func unicast_receive(state float64, stateQueue *[]float64) {
	_, _, numNodes, numFailures := Utils.FetchConfig()

	*stateQueue = append(*stateQueue, state)
	average := 0.0
	// if state > 0 {
	// 	fmt.Printf("\nReceived %f, system time is %s\n", state, time.Now().Format("Jan _2 15:04:05.000"))
	// 	fmt.Print(*stateQueue)
	// }
	total := 0.0

	for _, value := range *stateQueue {
		total += value
	}

	average = float64(total) / float64(len(*stateQueue))

	//wait for n-f messages and then update state
	if len(*stateQueue) >= numNodes-numFailures {
		fmt.Println(average)
	}

	// Unicast_send(average, 2)

	//*stateQueue = append(*stateQueue, msg.State)
	//fmt.Println(*stateQueue)
	// averageState := 0.0
	// if msg.State > 0 {
	//fmt.Printf("\nReceived %f, system time is %s\n", msg.State, time.Now().Format("Jan _2 15:04:05.000"))
	// *pmsg = append(*pmsg, msg)
	// fmt.Println(*stateQueue)
	//fmt.Println(*pmsg)
	//Wait until n-f messages received
	// sum := 0.00
	// l := len(*pmsg)
	//Suppose n-f = l
	// if l == 2 {
	// for i := 0; i < l; i++ {
	// 	sum += ((*stateQueue)[i])
	// }
	// updateState := (float64(sum)) / (float64(l))
	//R := msg.R + 1
	// averageState = updateState
	//averageState.R = R
	//fmt.Print("Update message is",averageState)

	// fmt.Println(averageState)
	// return &averageState
	// }
	// }
	// fmt.Println(averageState)
	// return &averageState
}

//func handleConnection(c net.Conn, stateQueue *[]float64, pmsg *[]Utils.Message) {
func handleConnection(c net.Conn, stateQueue *[]float64) {
	for {
		decoder := gob.NewDecoder(c) //initialize gob decoder
		//msg := new(Utils.Message)
		state := new(float64)

		_ = decoder.Decode(state)
		unicast_receive(*state, stateQueue)

		// encoder := gob.NewEncoder(c)
		// encoder.Encode(*averageState)
	}
	c.Close()
}

func StartServer(NodeNum string) {
	var stateQueue []float64
	p := &stateQueue
	//var msgs []Utils.Message
	//pmsg := &msgs
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
		go handleConnection(c, p)
	}
}

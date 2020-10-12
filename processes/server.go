package processes

import (
	"encoding/gob"
	"fmt"
	"net"
	"time"

	"../Utils"
)

func unicast_receive(msg Utils.Message , pmsg *[]Utils.Message, stateQueue *[]float64) {
	if msg.State > 0 {
		fmt.Printf("\nReceived %f, system time is %s\n", msg.State, time.Now().Format("Jan _2 15:04:05.000"))
		*stateQueue = append(*stateQueue, msg.State)
		*pmsg = append(*pmsg, msg)
		fmt.Println(*stateQueue)
		fmt.Println(*pmsg)
	}

}

func handleConnection(c net.Conn, stateQueue *[]float64, pmsg *[]Utils.Message) {
	for {
		decoder := gob.NewDecoder(c) //initialize gob decoder
		msg := new(Utils.Message)

		_ = decoder.Decode(msg)
		unicast_receive(*msg, pmsg, stateQueue)
	}
	c.Close()
}

func StartServer(NodeNum string) {
	var stateQueue []float64
	p := &stateQueue
	var msgs []Utils.Message
	pmsg := &msgs
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
		go handleConnection(c, p, pmsg)
	}
}

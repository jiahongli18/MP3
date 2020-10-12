package processes

import (
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
	"sync"

	"../Utils"
)

//Reads user command and sends the message with regards to destination and delay bounds
func Unicast_send(initmsg Utils.Message, n int) {
	for {
		for NodeNum := 1; NodeNum <= n; NodeNum++ {
			//find the associating host/port according to the user's desired destination #
			SNum := strconv.Itoa(NodeNum)
			ip, port, _ := Utils.FetchHostPort(SNum)
			min_delay, max_delay := Utils.FetchDelay()
			unicast_send(SNum, ip+":"+port, initmsg, min_delay, max_delay)
		}
	}

}

//Sends message to the destination process
func unicast_send(process string, destination string, initmsg Utils.Message, min_delay int, max_delay int) {
	//dial to the TCP server using net library
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		fmt.Println(err)
		return
	}

	encoder := gob.NewEncoder(conn)
	encoder.Encode(initmsg)
	//_ = encoder.Encode()

	//fmt.Printf("Sent %f to node %s, system time is %s\n", initmsg.State, process, time.Now().Format("Jan _2 15:04:05.000"))
	decoder := gob.NewDecoder(conn) //initialize gob decoder
	var updatemsg Utils.Message
	_ = decoder.Decode(&updatemsg)
	if updatemsg.State > 0 {
		fmt.Println("updatemsg is", updatemsg)
	}

	//set delay
	groupTest := new(sync.WaitGroup)
	go Utils.Delay(min_delay, max_delay, groupTest)

	//Wait group is used to block the execution of code in the main thread until all goroutines are complete and waitgroup counter is decremented to 0
	groupTest.Add(1)
	groupTest.Wait()
	// send to socket
	//fmt.Fprintf(conn, process + " " + state)
	// fmt.Fprintf(conn, process)

}

package processes

import (
	"../Utils"
	"encoding/gob"
	"fmt"
	"net"
	"strconv"
	"sync"

	//"bufio"
	//"os"
	//"strings"
	"time"
)

//Reads user command and sends the message with regards to destination and delay bounds
func Unicast_send(initstate float64, n int) {
	for NodeNum := 1;  NodeNum < n; NodeNum++{

		//find the associating host/port according to the user's desired destination #
		SNum := strconv.Itoa(NodeNum)
		ip, port := Utils.FetchHostPort(SNum)

		unicast_send(SNum, ip + ":" + port, initstate)

	}
}

//Sends message to the destination process
func unicast_send(process string, destination string, state float64) {
	//dial to the TCP server using net library
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		fmt.Println(err)
		return
	}

	encoder := gob.NewEncoder(conn)
	_ = encoder.Encode(state)

	fmt.Printf("Sent %f to node %s, system time is %s\n", state, process, time.Now().Format("Jan _2 15:04:05.000"))

	//set delay
	groupTest := new(sync.WaitGroup)

	//Wait group is used to block the execution of code in the main thread until all goroutines are complete and waitgroup counter is decremented to 0
	groupTest.Add(1)
	groupTest.Wait()
	// send to socket
	//fmt.Fprintf(conn, process + " " + state)
	fmt.Fprintf(conn,process)

}
package client

import (
	"encoding/gob"
	"fmt"
	"net"
	"os"
	"sync"
	"time"

	types "../structs"
	utils "../utils"
)

//The listen function waits on messages from the TCP server.
func Listen(c net.Conn, channel chan string) {
	for {
		decoder := gob.NewDecoder(c) //initialize gob decoder
		//Decode message struct and print it
		message := new(types.Message)
		_ = decoder.Decode(&message)

		if (*message == types.Message{"0", 0, false}) {
			//If the server terminates, the functions sends the exit signal through the channel to the main thread
			c.Close()
			os.Exit(0)
			channel <- "EXIT"
		}
	}
}

//Reads user command and sends the message with regards to destination and delay bounds
func Unicast(nodeNumber string) {

	initialValue := utils.FetchValue(nodeNumber)
	ports := utils.FetchPorts()
	//sends messages for r number of rounds
	for r := 0; r < 1; r++ {
		//find the associating host/port according to the user's desired destination #
		// address := utils.ReadStringFromConfig("host:port")
		minDelay := utils.ReadIntFromConfig("minDelay")
		maxDelay := utils.ReadIntFromConfig("maxDelay")

		for i := 0; i < len(ports); i++ {
			unicast_send(nodeNumber, "127.0.0.1:"+ports[i], initialValue, minDelay, maxDelay)
		}

	}
}

//Sends message to the destination process
func unicast_send(process string, destination string, message float64, minDelay int, maxDelay int) {
	//dial to the TCP server using net library
	conn, err := net.Dial("tcp", destination)
	if err != nil {
		fmt.Println(err)
		return
	}

	fmt.Printf("Sent %f to process %s, system time is %s\n", message, process, time.Now().Format("Jan _2 15:04:05.000"))

	//set delay
	groupTest := new(sync.WaitGroup)
	go utils.Delay(minDelay, maxDelay, groupTest)

	//Wait group is used to block the execution of code in the main thread until all goroutines are complete and waitgroup counter is decremented to 0
	groupTest.Add(1)
	groupTest.Wait()
	// send to socket
	encoder := gob.NewEncoder(conn)
	msg := types.Message{"0", message, false}

	encoder.Encode(msg)

}

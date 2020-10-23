package server

import (
	"encoding/gob"
	"fmt"
	"net"

	types "../structs"
	utils "../utils"
)

//Delivers the message received from the source process and prints out message, source, and time
// func unicastReceive(source string, message string) {
// }

func handleConnection(c net.Conn, processNum string, stateQueue *[]float64, numOfNodes int, numOfFaultyNodes int) {

	decoder := gob.NewDecoder(c) //initialize gob decoder
	message := new(types.Message)
	decoder.Decode(&message)

	// checkIfFaulty(processNum, message.Round)

	sum := 0.0
	if message.Value > 0.0 {
		*stateQueue = append(*stateQueue, message.Value)
		sum += message.Value
	}

	if len(*stateQueue) >= numOfNodes-numOfFaultyNodes {
		// fmt.Println(*stateQueue)
		fmt.Println("Average =", sum/float64(len(*stateQueue)))
		fmt.Println(message.Round)
	}

}

func StartServer(processNum string) {
	var stateQueue []float64
	numOfNodes := utils.ReadIntFromConfig("numOfNodes")
	numOfFaultyNodes := utils.ReadIntFromConfig("numOfFaultyNodes")
	_, port := utils.FetchHostPort(processNum)

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
		handleConnection(c, processNum, &stateQueue, numOfNodes, numOfFaultyNodes)
	}
}

// func checkIfFaulty(processNum string, round int) {
// 	faultyNodes := utils.FetchFaultyNodes()
// 	rand.Seed(time.Now().UnixNano())
// 	min := 0
// 	max := 10

// 	for i := 0; i < len(faultyNodes); i++ {
// 		if processNum == faultyNodes[i] {
// 			num := rand.Intn(max-min+1) + min
// 			if num <= 2 && round > 0 {
// 				fmt.Println(round)
// 				os.Exit(0)
// 			}
// 		}
// 	}
// }

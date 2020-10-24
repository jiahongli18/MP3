#MP3
Unfortunately, we are unable to complete the network design in time, specifically the construction of sending the round numbers between the controller and the nodes. Here we write out the steps we plan to start the simulation and the logic behind future implementation.

##Usage
Go to the controller folder, in one terminal, start the controller server and enter 1 as an argument

```bash
go run controller.go 1
```

In another terminal, start all the nodes as main programs running in the background. 
```bash
bash script.sh
```

##Structure and Design
* The steps in the algorithm will be displayed like this:
1) all nodes send their value to each other once
2) take average of all values reecived, then send that to the central server
3) wait for reply from server on whether to stop or keep going
4) if the server says to stop, then stop and end execution. Else, run step 1 again with new value.

* We intend to follow Professor Tseng's idea and build our design based on the integration of MP1 and MP2. Like MP2, we make a controller, which is a central server that all the nodes dial to. The controller has the information of number of faulty nodes and can prompt all the nodes to exit once the approximate consensus is achieved. 

In the controller, we have two goroutines: `startServer()` and `exit()`
* `startServer()` creates the TCP connection and listen on provided port for requests. We uses a map to store the channel as the associating value. For each connection, we call a goroutine to handle the communication so that the server can support handling multiple concurrent clients. The goroutine will communicate with the ndoes using decoder. 

* `exit()` is used to determine when the approximate consensus is achieved. If the state difference between each nodes is not above 0.001, then a channel is used to communicate this information with the main thread. Then the main thread sends the signal to all other TCP channels to terminates those, and finally terminates itself.

* Like MP2, the nodes in main.go act both like servers and clients which communicate with each other. We start two goroutines within main.go The first one is called `startServer()` and the second is `Listen()`

* `startServer()` creates the TCP connection and listen on provided port for requests. It gets the number of faulty ndoes and the number of nodes, and it makes an array called statequeue which stores the floats received from n-f nodes. For each connection, we call a goroutine to handle the communication so that the each node can decode the message received from other nodes and calculate the average of the n-f states it receives this round. 

* `Listen()` waits on messages from the controller and decodes it. Once it receives exit message, each node will terminate. 

* `Unicast()` is a function that fetches the initialValue, minDelay, and maxDelay from the JSON file and sends to all the nodes including itself. 

* `Initialize()` we keep the function Initialize() we have in MP1 to create the dial connections between each pair of processes in the config file. We put this process in initialize() in main.go. The idea is to try alive ports and dial between the process itself and every other process in the config file. This connection is attempted every two seconds until a successful connection is made(we coded a delay here because we didn't want the program to spam the connections too fast).

* In utils.go, we abstract helper functions for each nodes to fetch host and ports, nodeNumbers, faultyNodeNumbers, mindelay, and maxdelay.

* We have three structs as Config, Node, and Message with explicit fields listed below. 
type Config struct {
	MinDelay         int      `json:"minDelay"`
	MaxDelay         int      `json:"maxDelay"`
	Host             string   `json:"host"`
	DefaultPort      string   `json:"defaultPort"`
	NumOfFaultyNodes int      `json:"numOfFaultyNodes"`
	FaultyNodes      []string `json:"faultyNodes"`
	Nodes            []Node   `json:"nodes"`
}

type Node struct {
	Number string  `json:"number"`
	Port   string  `json:"port"`
	Value  float64 `json:"value"`
}

type Message struct {
	Round  string
	Value  float64
	IsExit bool
}

## Chanllenges and Bottleneck
We cannot figure out how to increment the round number correctly. Once the round number is added into the message, the goroutines seem to go wrong for unknown reasons. We have to halt our attempt and fix it later. 


## Resources
* [TCP Concurrent Server](https://www.linode.com/docs/development/go/developing-udp-and-tcp-clients-and-servers-in-go/)
* [Gob](https://golang.org/pkg/encoding/gob/)
* [Parsing JSON](https://tutorialedge.net/golang/parsing-json-with-golang/)
* Professor Tseng's design 

## Authors
* Jiahong Li
* Zheng Zhou
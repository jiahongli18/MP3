package Utils

import (
	"bufio"
	"log"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"sync"
	"time"
)

//Fetches all the ports
func FetchPorts() []string {
	line := 0
	f, err := os.Open("./config.txt")
	var ports []string
	if err != nil {
		log.Fatal(err)
	}

	defer f.Close()
	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if line != 0 {
			port := strings.Split(scanner.Text(), " ")[2]
			ports = append(ports, port)
		}
		line = line + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}
	return ports
}

//parses config.txt and returns ip and host
func FetchHostPort(destination string) (string, string, float64) {
	line := 0
	f, err := os.Open("./config.txt")
	if err != nil {

		log.Fatal(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)

	for scanner.Scan() {
		if line != 0 {
			process := strings.Split(scanner.Text(), " ")[0]
			ip := strings.Split(scanner.Text(), " ")[1]
			port := strings.Split(scanner.Text(), " ")[2]
			stateStr := strings.Split(scanner.Text(), " ")[3]
			state, _ := strconv.ParseFloat(stateStr, 64)
			if process == destination {
				//fmt.Println(ip, port)
				return ip, port, state

			}
		}

		line = line + 1
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	return "", "nn", 0.64
}

//parses the min and max delays from the config file
func FetchConfig() (int, int, int, int) {
	f, err := os.Open("./config.txt")
	if err != nil {
		log.Fatal(err)
	}

	scanner := bufio.NewScanner(f)
	scanner.Scan()
	delays := strings.Fields(scanner.Text())
	min_delay, _ := strconv.Atoi(delays[0])
	max_delay, _ := strconv.Atoi(delays[1])
	numNodes, _ := strconv.Atoi(delays[2])
	numFailures, _ := strconv.Atoi(delays[3])
	f.Close()
	return min_delay, max_delay, numNodes, numFailures
}

//Simulate network delay by adding an extra layer before sending the message via the TCP channel
func Delay(min int, max int, wg *sync.WaitGroup) {
	num := rand.Intn(max-min) + min
	time.Sleep(time.Duration(num) * time.Millisecond)

	//decrement value of waitgroup and relay the flow of execution back to main
	wg.Done()
}

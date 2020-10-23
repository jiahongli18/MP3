package utils

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"math/rand"
	"os"
	"sync"
	"time"

	types "../structs"
)

//reads strings from the config file
func ReadStringFromConfig(desiredInfo string) (address string) {
	// Open  jsonFile
	jsonFile, err := os.Open("/Users/jiahongli/Desktop/DistributedSystems/MP3/config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	config := types.Config{}
	json.Unmarshal(byteValue, &config)

	//return desired info
	if desiredInfo == "host:port" {
		return config.Host + ":" + config.DefaultPort
	} else if desiredInfo == "port" {
		return config.DefaultPort
	} else {
		return ""
	}
}

//reads ints from the config file
func ReadIntFromConfig(desiredInfo string) (data int) {
	// Open  jsonFile
	jsonFile, err := os.Open("/Users/jiahongli/Desktop/DistributedSystems/MP3/config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	config := types.Config{}
	json.Unmarshal(byteValue, &config)

	//return desired info
	if desiredInfo == "numOfNodes" {
		return len(config.Nodes)
	} else if desiredInfo == "numOfFaultyNodes" {
		return config.NumOfFaultyNodes
	} else if desiredInfo == "minDelay" {
		return config.MinDelay
	} else if desiredInfo == "maxDelay" {
		return config.MaxDelay
	} else {
		return 0
	}
}

func FetchPorts() []string {
	var ports []string

	// Open  jsonFile
	jsonFile, err := os.Open("/Users/jiahongli/Desktop/DistributedSystems/MP3/config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	config := types.Config{}
	json.Unmarshal(byteValue, &config)

	for i := 0; i < len(config.Nodes); i++ {
		ports = append(ports, config.Nodes[i].Port)
	}

	return ports
}

func FetchFaultyNodes() []string {
	// Open  jsonFile
	jsonFile, err := os.Open("/Users/jiahongli/Desktop/DistributedSystems/MP3/config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	config := types.Config{}
	json.Unmarshal(byteValue, &config)

	return config.FaultyNodes
}

//parses config.json and returns host and port
func FetchHostPort(processNum string) (string, string) {
	// Open  jsonFile
	jsonFile, err := os.Open("/Users/jiahongli/Desktop/DistributedSystems/MP3/config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	config := types.Config{}
	json.Unmarshal(byteValue, &config)

	for i := 0; i < len(config.Nodes); i++ {
		if config.Nodes[i].Number == processNum {
			return config.Host, config.Nodes[i].Port
		}
	}

	return "null", "null"
}

func FetchValue(processNum string) float64 {
	// Open  jsonFile
	jsonFile, err := os.Open("/Users/jiahongli/Desktop/DistributedSystems/MP3/config.json")
	if err != nil {
		fmt.Println(err)
	}
	defer jsonFile.Close()

	// read our opened xmlFile as a byte array.
	byteValue, _ := ioutil.ReadAll(jsonFile)

	// we initialize our Users array
	config := types.Config{}
	json.Unmarshal(byteValue, &config)

	for i := 0; i < len(config.Nodes); i++ {
		if config.Nodes[i].Number == processNum {
			return config.Nodes[i].Value
		}
	}

	return 0.0
}

func Delay(min int, max int, wg *sync.WaitGroup) {
	num := rand.Intn(max-min) + min
	time.Sleep(time.Duration(num) * time.Millisecond)

	//decrement value of waitgroup and relay the flow of execution back to main
	wg.Done()
}

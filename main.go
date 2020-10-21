package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"

	types "./Utils"
)

func main() {
	readConfig()
}

func readConfig() {
	// Open  jsonFile
	jsonFile, err := os.Open("config.json")
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
		fmt.Println(config.Nodes[i].Value)
	}
}

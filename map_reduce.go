package main

import (
	"os"
	"strconv"
	"strings"
	"time"
)

var HOSTS = []string{"host1", "host2"}

func getMapInput(host string) *os.File {
	data, _ := os.Open(host + "/latencies.txt")
	return data
}

func getReduceInput() map[string]string {
	files, _ := os.ReadDir("./map_results")
	inputs := make(map[string]string)
	for _, f := range files {
		fName := strings.Split(f.Name(), ".")
		key := fName[0]
		content, _ := os.ReadFile("./map_results/" + f.Name())
		value := strconv.Itoa(len(string(content)) / 2)
		inputs[key] = value
	}
	return inputs
}

func doMap(host string) {
	mapInput := getMapInput(host)
	Map(host, mapInput)
}
func main() {
	// do map op for both hosts
	// this simulates a distributed systems
	go doMap(HOSTS[0])
	go doMap(HOSTS[1])
	time.Sleep(time.Second)
	Shuffle(HOSTS)
	reduceInput := getReduceInput()
	Reduce(reduceInput)
}

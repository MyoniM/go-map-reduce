package main

import (
	"bufio"
	"os"
	"strconv"
)

func Map(host string, mapInput *os.File) {
	u, _ := os.OpenFile(host+"/map_results/under_1_second.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer u.Close()
	a, _ := os.OpenFile(host+"/map_results/above_1_second.txt", os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
	defer a.Close()
	scanner := bufio.NewScanner(mapInput)
	for scanner.Scan() {
		latency := scanner.Text()
		l, _ := strconv.Atoi(latency)
		if l < 1000 {
			u.WriteString("1\n")
		} else {
			a.WriteString("1\n")
		}
	}
}

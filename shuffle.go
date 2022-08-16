package main

import (
	"os"
)

// this function creates intermediate result of the map function
func Shuffle(HOSTS []string) {
	for _, host := range HOSTS {
		files, _ := os.ReadDir(host + "/map_results")
		for _, f := range files {
			content, _ := os.ReadFile(host + "/map_results/" + f.Name())
			u, _ := os.OpenFile("./map_results/"+f.Name(), os.O_APPEND|os.O_WRONLY|os.O_CREATE, 0644)
			defer u.Close()
			u.WriteString(string(content))
		}
	}
}

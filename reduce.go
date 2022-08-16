package main

import "os"

func Reduce(reduceInput map[string]string) {
	r, _ := os.OpenFile("./reduce_results/result.txt", os.O_CREATE, 0644)
	defer r.Close()
	for key, value := range reduceInput {
		r.WriteString(key + " " + value + "\n")
	}
}

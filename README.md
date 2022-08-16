## go-map-reduce 

This is to simulate a basic Map Reduce procedure on a distributed system.
```host1``` and ```host2``` are machines with their request latencies log.
The final result of the Map Reduce will be a file with count of latencies above one second and below one second.

```go
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
```

Map function is called on both hosts. Here, the call is concurrent because in a distributed systems, every machine would call the Map function simultaneously.
After ```doMap``` the ```Shuffle``` function is called to create the intermediate arguments for the reduce function.
Finally, the ```Reduce``` function is called and generates the required result inside ```reduce_results/result.txt``` file.

#! /bin/bash

# clean up from previous run
rm -f host1/map_results/*.txt
rm -f host2/map_results/*.txt
rm -f map_results/*.txt
rm -f reduce_results/result.txt

# run map_reduce
# map_reduce.go will call map for both hosts concurrently  
go run map_reduce.go map.go shuffle.go reduce.go 

# view
cat reduce_results/result.txt
package main

import (
	"fmt"
	"log"
	"net/rpc"
	"sync"
)

type ChunkArgs struct {
	Input []string
	WorkerId int
}

type ChunkResult struct {
	L, R int
}

func main() {
	var wg sync.WaitGroup
	client, err := rpc.Dial("tcp", "localhost:8080")

	if err != nil {
		fmt.Println("Could not connect rpc")
		return
	}

	defer client.Close()

	const NUM_OF_WORKERS = 3
	inputFilepath := "/Users/fnuworsu/Distributed Systems/mapReduce/client/input.txt"
	var input []string

	err = client.Call("Coordinator.GetInput", inputFilepath, &input)

	if err != nil {
		log.Fatal(err)
		return
	}

	for workerId := 0; workerId < NUM_OF_WORKERS; workerId++ {
		wg.Add(1)

		go func(id int) {
			defer wg.Done()

			args := &ChunkArgs{Input: input, WorkerId: workerId}
			var res ChunkResult
			outputFilePath := "/Users/fnuworsu/Distributed Systems/mapReduce/client/result.txt"
			mappedOutputs := make(chan []map[string]int, 1)

			err := client.Call("Coordinator.Chunk", args, &res)
			if err != nil {
				log.Fatal(err)
				return
			}
			
			err = client.Call("Coordinator.PhaseOneWorker", input[res.L : res.R], mappedOutputs)
			if err != nil {
				log.Fatal(err)
				return
			}

			err = client.Call("Coordinator.PhaseTwoWorker", <-mappedOutputs, outputFilePath)
			if err != nil {
				log.Fatal(err)
				return
			}
		}(workerId)

		wg.Wait()
		fmt.Println("All workers completed!")
	}
}
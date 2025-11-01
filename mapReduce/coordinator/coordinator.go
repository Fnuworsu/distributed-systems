package coordinator

import (
	"bufio"
	"fmt"
	"log"
	"mapReduce/worker"
	"os"
	"strings"
	"sync"
)
const NUM_OF_WORKERS = 3

type Coordinator struct{}

type ChunkArgs struct {
	Input []string
	WorkerId int
}

type ChunkResult struct {
	L, R int
}

func (coord *Coordinator) GetInput(filepath string, result *[]string) error {
	file, err := os.Open(filepath)

	if err != nil {
		return err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)

	for scanner.Scan() {
		line := scanner.Text()
		*result = append(*result, strings.TrimSpace(line))
	}

	err = scanner.Err()

	if err != nil {
		return err
	}

	return nil
}

func (coord *Coordinator) Chunk(args ChunkArgs, result *ChunkResult) error {
	inputLen := len(args.Input)
	chunk_size := inputLen / NUM_OF_WORKERS
	start := args.WorkerId * chunk_size
	end := 0

	if args.WorkerId == NUM_OF_WORKERS - 1 {
		end = inputLen
	} else {
		end = start + chunk_size
	}

	*result = ChunkResult{L: start, R: end}
	return nil
}

func (coord *Coordinator) PhaseOneWorker(chunkedInput []string, mappedOutputs chan []map[string]int) error {
	mapped := make(chan map[string]int, len(chunkedInput))
	var results []map[string]int
	var wg sync.WaitGroup

	for _, line := range chunkedInput {
		wg.Add(1)
		go func(l string) {
			defer wg.Done()
			worker.Mapper(l, mapped)
		}(line)
	}

	go func()  {
		wg.Wait()
		close(mapped)
	}()

	for i := 0; i < len(chunkedInput); i++ {
		results = append(results, <-mapped)
	}

	mappedOutputs <- results
	return nil
}

func (coord *Coordinator) PhaseTwoWorker(mappedOutputs chan []map[string]int, filePath string) error {
	shuffled := make(chan map[string][]int, 1)
	reduced := make(chan map[string]int, 1)
	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		worker.Shuffle(<-mappedOutputs, shuffled)
	}()

	wg.Add(1)
	go func() {
		defer wg.Done()
		worker.Reducer(<-shuffled, reduced)
	}()

	wg.Wait()

	file, err := os.Create(filePath)

	if err != nil {
		log.Fatal(err)
	}

	defer file.Close()

	for key, val := range <-reduced {
		line := fmt.Sprintf("(%s, %d)", key, val)
		_, err := file.WriteString(line + "\n")

		if err != nil {
			log.Fatal(err)
		}
	}

	return nil
}

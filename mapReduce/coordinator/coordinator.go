package main

import (
	"bufio"
	"os"
	"strings"
)
const NUM_OF_WORKERS = 3

type Coordinator struct{}

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

func (coord *Coordinator) Chunk(input []string, workerId int, l *int, r *int) error {
	inputLen := len(input)
	chunk_size := inputLen / NUM_OF_WORKERS
	start := workerId * chunk_size
	end := 0

	if workerId == NUM_OF_WORKERS - 1 {
		end = inputLen
	} else {
		end = start + chunk_size
	}

	*l, *r = start, end
	return nil
}



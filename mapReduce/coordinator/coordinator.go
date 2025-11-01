package coordinator

import (
	"bufio"
	"os"
	"strings"
)

const NUM_OF_WORKERS = 3

func getInput(filepath string) ([]string, error) {
	file, err := os.Open(filepath)

	if err != nil {
		return nil, err
	}

	defer file.Close()

	scanner := bufio.NewScanner(file)
	var res []string

	for scanner.Scan() {
		line := scanner.Text()
		res = append(res, strings.TrimSpace(line))
	}

	err = scanner.Err()

	if err != nil {
		return nil, err
	}

	return res, nil
}

func chunk(input []string, workerId int) (l,r int) {
	/*
	 * number of workers = w
	 * len of input = n
	 * chunk_size = w / n
	*/
	inputLen := len(input)
	chunk_size := inputLen / NUM_OF_WORKERS
	start := workerId * chunk_size
	end := 0

	if workerId == NUM_OF_WORKERS - 1 {
		end = inputLen
	} else {
		end = start + chunk_size
	}

	return start, end
}

// TODO: Implement worker function
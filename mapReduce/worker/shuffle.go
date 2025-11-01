package worker

func Shuffle(mappedOutputs []map[string]int, ch chan map[string][]int) {
	shuffled := make(map[string][]int)

	for _, mapped := range mappedOutputs {
		for word, count := range mapped {
			shuffled[word] = append(shuffled[word], count)
		}
	}

	ch <- shuffled
}
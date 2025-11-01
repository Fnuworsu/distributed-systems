package worker

type Pair struct {
	Key string
	Val int
}

func Shuffle(mappedOutputs []map[string]int, outCh chan <- *Pair) {
	for _, mapped := range mappedOutputs {
		for word, count := range mapped {
			outCh <- &Pair{Key: word, Val: count}
		}
	}
}
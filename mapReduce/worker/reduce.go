package worker

func Reducer(shuffled map[string][]int, reduced *map[string]int) {
	sum := func(arr []int) int {
		res := 0

		for _, n := range arr {
			res += n
		}

		return res
	}

	for word, counts := range shuffled {
		(*reduced)[word] = sum(counts)
	}
}

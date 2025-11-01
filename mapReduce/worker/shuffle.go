package worker

import "sort"

func shuffle(emitted []*Emit) []*Emit {
	sort.Slice(emitted, func(i, j int) bool {
		return emitted[i].Val < emitted[j].Val
	})

	return emitted
}
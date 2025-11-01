package worker

import "strings"

func Mapper(line string) map[string]int {
	mapped := make(map[string]int)

	for _,word := range strings.Split(line, "") {
		_,ok := mapped[word]

		if ok {
			mapped[word]++
		} else {
			mapped[word] = 1
		}
	}

	return mapped
}
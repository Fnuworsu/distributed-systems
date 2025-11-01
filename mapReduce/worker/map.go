package worker

func mapper(word string) map[rune]int {
	mapp := make(map[rune]int)

	for _,char := range word {
		_, ok := mapp[char]

		if ok {
			mapp[char]++
		} else {
			mapp[char] = 1
		}
	}

	return mapp
}
package worker

type Emit struct {
	Key rune
	Val int
}

func reducer(mapp []map[rune]int) []*Emit {
	resMap := make(map[rune]int)

	for _, item := range mapp {
		for key,val := range item {
			_,ok := resMap[key]

			if ok {
				resMap[key] += val
			} else {
				resMap[key] = val
			}
		}
	}

	var resEmit []*Emit

	for key,val := range resMap {
		resEmit = append(resEmit, &Emit{Key: key, Val: val})
	}

	return resEmit
}
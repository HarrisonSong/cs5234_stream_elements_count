package main

func processStream(stream []int) (countMap map[int]int) {
	countMap = map[int]int{}
	for i := 0; i < len(stream); i++ {
		countMap[stream[i]]++
	}

	return
}

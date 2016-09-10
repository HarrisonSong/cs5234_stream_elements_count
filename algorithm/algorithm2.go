package algorithm

import (
	"sort"

	"github.com/HarrisonSong/cs5234_stream_elements_count/hashing"
)

type Algorithm2 struct {
	Algorithm
}

func (a *Algorithm2) Query(x int) int {
	results := make([]int, len(a.counters))
	for i := 0; i < a.HashFunctions.GetHashingFunctionNumber(); i++ {
		hashResult := a.HashFunctions.GethashingFunction(i)(x)
		results[i] = a.counters[i][hashResult] - a.counters[i][neighbor(hashResult)]
	}

	return sort.IntSlice(results)[len(results)/2]
}

func Algorithm2Initializer(hf hashing.HashingFamily, s []int, B int) Algorithm2 {
	return Algorithm2{AlgorithmInitializer(hf, s, B)}
}

func neighbor(index int) int {
	if index%2 == 0 {
		return index + 1
	} else {
		return index - 1
	}
}

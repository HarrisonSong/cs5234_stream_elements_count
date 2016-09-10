package algorithm

import (
	"sort"

	"github.com/HarrisonSong/cs5234_stream_elements_count/hashing"
)

type Algorithm1 struct {
	Algorithm
}

func (a *Algorithm1) Query(x int) int {
	results := make([]int, len(a.counters))
	for i := 0; i < a.HashFunctions.GetHashingFunctionNumber(); i++ {
		results[i] = a.counters[i][a.HashFunctions.GethashingFunction(i)(x)]
	}

	return sort.IntSlice(results)[len(results)/2]
}

func Algorithm1Initializer(hf hashing.HashingFamily, s []int, B int) Algorithm1 {
	return Algorithm1{AlgorithmInitializer(hf, s, B)}
}

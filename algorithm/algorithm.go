package algorithm

import "github.com/HarrisonSong/cs5234_stream_elements_count/hashing"

type Algorithm struct {
	HashFunctions hashing.HashingFamily
	counters      [][]int
}

func AlgorithmInitializer(hf hashing.HashingFamily, s []int, B int) Algorithm {
	c := [][]int{}
	for i := 0; i < hf.GetHashingFunctionNumber(); i++ {
		c = append(c, make([]int, B))
		hashFunction := hf.GethashingFunction(i)
		for j := 0; j < len(s); j++ {
			c[i][hashFunction(s[j])]++
		}
	}

	return Algorithm{hf, c}
}

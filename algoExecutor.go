package main

import (
	"math"
	"math/rand"
	"time"

	"github.com/HarrisonSong/cs5234_stream_elements_count/algorithm"
	"github.com/HarrisonSong/cs5234_stream_elements_count/hashing"
	"github.com/HarrisonSong/cs5234_stream_elements_count/stream"
)

func RunAlgo(N, M, B, A, alg, algoRepeatTime, totalTrials int, errorRate float64, streamType string) float64 {
	str := []int{}
	if streamType == "u" {
		str = stream.GenerateUniformIntegerStream(N, M)
	} else if streamType == "n" {
		str = stream.GenerateNormalIntegerStream(N, M)
	} else {
		str = stream.GenerateExponentialIntegerStream(N, M)
	}
	countMap := processStream(str)
	IHF := hashing.IntegerHashingFamilyInitializer(A, B)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	correctness := 0
	if alg == 1 {
		algorithm1 := algorithm.Algorithm1Initializer(IHF, str, B)
		for i := 0; i < totalTrials; i++ {
			n := r.Intn(M)
			if isCorrect(algorithm1.Query(n), countMap[n], float64(N)/float64(M), errorRate) {
				correctness++
			}
		}
	} else {
		algorithm2 := algorithm.Algorithm2Initializer(IHF, str, B)
		for i := 0; i < totalTrials; i++ {
			n := r.Intn(M)
			if isCorrect(algorithm2.Query(n), countMap[n], float64(N)/float64(M), errorRate) {
				correctness++
			}
		}
	}
	accuracy := (float64(correctness) / float64(totalTrials)) * 100
	return accuracy
}

func isCorrect(realValue, expectedValue int, expectation, errorRate float64) bool {
	return math.Abs(float64(realValue-expectedValue))/float64(expectedValue) <= errorRate
}

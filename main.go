package main

import (
	"flag"
	"io/ioutil"
	"log"
	"math"
	"math/rand"
	"strconv"
	"time"

	"github.com/HarrisonSong/cs5234_stream_elements_count/algorithm"
	"github.com/HarrisonSong/cs5234_stream_elements_count/hashing"
	"github.com/HarrisonSong/cs5234_stream_elements_count/stream"
)

func main() {
	N := flag.Int("N", 100000, "total number of elements in stream.")
	M := flag.Int("M", 10000, "range of elements in stream. limited in 2^32")
	B := flag.Int("B", 1000, "range of hashing")
	A := flag.Int("A", 50, "number of hash functions")
	alg := flag.Int("alg", 1, "algorithm to be run")
	algoRepeatTime := flag.Int("repreat", 10, "algorithm repeat time")
	totalTrials := flag.Int("times", 10000, "total trial times")
	errorRate := flag.Float64("erate", 0.05, "max acceptable error rate")
	flag.Parse()

	sum := 0.0
	for times := 0; times < *algoRepeatTime; times++ {
		sum += runAlgo(*N, *M, *B, *A, *alg, *algoRepeatTime, *totalTrials, *errorRate)
	}

	writeToFile("results/A_"+strconv.Itoa(*A)+"_B_"+strconv.Itoa(*B)+".csv", strconv.FormatFloat(sum/float64(*algoRepeatTime), 'f', -1, 64))
	log.Printf("final accuracy is %f\n", sum/float64(*algoRepeatTime))
}

func runAlgo(N, M, B, A, alg, algoRepeatTime, totalTrials int, errorRate float64) float64 {
	stream := stream.GenerateIntegerStream(N, M)
	countMap := processStream(stream)
	IHF := hashing.IntegerHashingFamilyInitializer(A, B)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	correctness := 0
	if alg == 1 {
		//startAlgorithm := time.Now()
		algorithm1 := algorithm.Algorithm1Initializer(IHF, stream, B)
		//log.Printf("setup hash algorithm 1 time: %s\n", time.Since(startAlgorithm))
		//startRuningAlgorithm := time.Now()
		for i := 0; i < totalTrials; i++ {
			n := r.Intn(M)
			if isCorrect(algorithm1.Query(n), countMap[n], float64(N)/float64(M), errorRate) {
				correctness++
			}
		}
		//log.Printf("runing algorithm 1 time: %s\n", time.Since(startRuningAlgorithm))
	} else {
		//startAlgorithm := time.Now()
		algorithm2 := algorithm.Algorithm2Initializer(IHF, stream, B)
		//log.Printf("setup hash algorithm 1 time: %s\n", time.Since(startAlgorithm))
		//startRuningAlgorithm := time.Now()
		for i := 0; i < totalTrials; i++ {
			n := r.Intn(M)
			if isCorrect(algorithm2.Query(n), countMap[n], float64(N)/float64(M), errorRate) {
				correctness++
			}
		}
		//log.Printf("runing algorithm 2 time: %s\n", time.Since(startRuningAlgorithm))
	}
	accuracy := (float64(correctness) / float64(totalTrials)) * 100
	log.Printf("accuracy: %f\n", accuracy)
	return accuracy
}

func isCorrect(realValue, expectedValue int, expectation, errorRate float64) bool {
	return math.Abs(float64(realValue-expectedValue))/float64(expectedValue) <= errorRate
}

func writeToFile(fileName string, content string) {
	err := ioutil.WriteFile(fileName, []byte(content), 0644)
	if err != nil {
		log.Println(err)
	}
}

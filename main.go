package main

import (
	"flag"
	"log"
	"math/rand"
	"time"

	"github.com/HarrisonSong/cs5234_stream_elements_count/algorithm"
	"github.com/HarrisonSong/cs5234_stream_elements_count/hashing"
	"github.com/HarrisonSong/cs5234_stream_elements_count/stream"
)

func main() {
	N := flag.Int("N", 10000, "total number of elements in stream")
	M := flag.Int("M", 1000, "range of elements in stream")
	B := flag.Int("B", 1000, "range of hashing")
	A := flag.Int("A", 100, "number of hash functions")
	alg := flag.Int("alg", 1, "algorithm to be run")
	totalTrials := flag.Int("times", 1000, "total trial times")
	flag.Parse()

	startStreaming := time.Now()
	stream := stream.GenerateIntegerStream(*N, *M)
	log.Printf("streaming time: %s\n", time.Since(startStreaming))

	countMap := processStream(stream)
	log.Printf("total different number %d\n", len(countMap))

	startHashing := time.Now()
	IHF := hashing.IntegerHashingFamilyInitializer(*A, *B)
	log.Printf("hashing time: %s\n", time.Since(startHashing))

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	if *alg == 1 {
		startAlgorithm := time.Now()
		algorithm1 := algorithm.Algorithm1Initializer(IHF, stream, *B)
		log.Printf("setup algorithm time: %s\n", time.Since(startAlgorithm))

		correctness := 0
		log.Println("Algorithm 1:")
		startRuningAlgorithm := time.Now()
		for i := 0; i < *totalTrials; i++ {
			n := r.Int()
			if algorithm1.Query(n) == countMap[n] {
				correctness++
			}
		}
		log.Printf("runing algorithm 1 time: %s\n", time.Since(startRuningAlgorithm))
		log.Printf("algorithm 1 accuracy: %f\n", (float64(correctness)/float64(*totalTrials))*100)
	} else {
		startAlgorithm := time.Now()
		algorithm2 := algorithm.Algorithm2Initializer(IHF, stream, *B)
		log.Printf("setup algorithm time: %s\n", time.Since(startAlgorithm))

		correctness := 0
		log.Println("Algorithm 2:")
		startRuningAlgorithm := time.Now()
		for i := 0; i < *totalTrials; i++ {
			n := r.Int()
			if algorithm2.Query(n) == countMap[n] {
				correctness++
			}
		}
		log.Printf("runing algorithm 2 time: %s\n", time.Since(startRuningAlgorithm))
		log.Printf("algorithm 2 accuracy: %f\n", (float64(correctness)/float64(*totalTrials))*100)
	}
}

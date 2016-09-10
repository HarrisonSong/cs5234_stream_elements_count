package main

import (
	"log"
	"math/rand"
	"time"

	"github.com/HarrisonSong/cs5234_stream_elements_count/algorithm"
	"github.com/HarrisonSong/cs5234_stream_elements_count/hashing"
	"github.com/HarrisonSong/cs5234_stream_elements_count/stream"
)

func main() {
	N := 100000
	M := 10000
	B := 1000
	A := 100

	startStreaming := time.Now()
	stream := stream.GenerateIntegerStream(N, M)
	log.Printf("streaming time: %s\n", time.Since(startStreaming))

	startHashing := time.Now()
	IHF := hashing.IntegerHashingFamilyInitializer(A, B)
	log.Printf("hashing time: %s\n", time.Since(startHashing))

	startAlgorithm := time.Now()
	algorithm1 := algorithm.Algorithm1Initializer(IHF, stream, B)
	algorithm2 := algorithm.Algorithm2Initializer(IHF, stream, B)
	log.Printf("setup algorithm time: %s\n", time.Since(startAlgorithm))

	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	log.Println("Algorithm 1:")
	startRuningAlgorithm := time.Now()
	for i := 0; i < 10; i++ {
		n := r.Int()
		log.Printf("number %d; query result: %d\n", n, algorithm1.Query(n))
	}
	log.Printf("runing algorithm 1 time: %s\n", time.Since(startRuningAlgorithm))

	log.Println("Algorithm 2:")
	startRuningAlgorithm = time.Now()
	for i := 0; i < 10; i++ {
		n := r.Int()
		log.Printf("number %d; query result: %d\n", n, algorithm2.Query(n))
	}
	log.Printf("runing algorithm 2 time: %s\n", time.Since(startRuningAlgorithm))
}

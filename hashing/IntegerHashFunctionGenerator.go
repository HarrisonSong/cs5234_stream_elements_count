package hashing

import (
	"math"
	"math/rand"
	"time"
)

var (
	// Commonly used Mersenne Prime 2^61 - 1
	prime = math.Pow(2, 61) - 1
)

type IntegerHashingGenerator struct {
	HashingSet []func(int) int
}

func IntegerHashingGeneratorInitializer(A, B int) IntegerHashingGenerator {
	IHG := IntegerHashingGenerator{make([]func(int) int, A)}

	// randomely select A sets of a and b and fill the HashingSet
	seedA := rand.NewSource(time.Now().UnixNano())
	randA := rand.New(seedA)
	for i := 0; i < A; i++ {
		IHG.HashingSet[i] = baseHashingGenerator(randA.Int(), randA.Int(), B)
	}

	return IHG
}

func (hs *IntegerHashingGenerator) GethashingFunction() func(int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// randomly select on hash function in HashingSet
	return hs.HashingSet[r.Intn(len(hs.HashingSet))]
}

func baseHashingGenerator(a, b, B int) func(int) int {
	return func(x int) int {
		return int(math.Mod(math.Mod(a*x+b, prime), B))
	}
}

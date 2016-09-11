package hashing

import (
	"math"
	"math/big"
	"math/rand"
	"time"
)

var (
	// Commonly used Mersenne Prime 2^61 - 1
	prime = int64(math.Pow(2, 61)) - 1
)

type IntegerHashingFamily struct {
	hashingSet []func(int) int
}

func IntegerHashingFamilyInitializer(A, B int) IntegerHashingFamily {
	IHG := IntegerHashingFamily{make([]func(int) int, A)}

	// randomly select A sets of a and b and fill the hashingSet
	seedA := rand.NewSource(time.Now().UnixNano())
	randA := rand.New(seedA)
	seedB := rand.NewSource(time.Now().UnixNano())
	randB := rand.New(seedB)
	for i := 0; i < int(A); i++ {
		IHG.hashingSet[i] = baseHashingGenerator(randA.Int63n(prime), randB.Int63n(prime+1)-1, B)
	}

	return IHG
}

func (hs IntegerHashingFamily) GethashingFunction(i int) func(int) int {
	return hs.hashingSet[i]
}

func (hs IntegerHashingFamily) GetHashingFunctionNumber() int {
	return len(hs.hashingSet)
}

func baseHashingGenerator(a, b int64, B int) func(int) int {
	return func(x int) int {
		result := big.NewInt(int64(x))
		result.Mul(result, big.NewInt(a)).Mod(result, big.NewInt(prime))
		return int((result.Int64() + b) % prime % int64(B))
	}
}

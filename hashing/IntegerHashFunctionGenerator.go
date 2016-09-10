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
	HashingSet []func(int) int
}

func IntegerHashingFamilyInitializer(A, B int) IntegerHashingFamily {
	IHG := IntegerHashingFamily{make([]func(int) int, A)}

	// randomely select A sets of a and b and fill the HashingSet
	seedA := rand.NewSource(time.Now().UnixNano())
	randA := rand.New(seedA)
	for i := 0; i < int(A); i++ {
		IHG.HashingSet[i] = baseHashingGenerator(randA.Int63n(prime), randA.Int63(), B)
	}

	return IHG
}

func (hs IntegerHashingFamily) GethashingFunction(i int) func(int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	// randomly select on hash function in HashingSet
	return hs.HashingSet[r.Intn(len(hs.HashingSet))]
}

func (hs IntegerHashingFamily) GetHashingFunctionNumber() int {
	return len(hs.HashingSet)
}

func baseHashingGenerator(a, b int64, B int) func(int) int {
	return func(x int) int {
		bA := big.NewInt(a)
		bB := big.NewInt(b)
		bPrime := big.NewInt(prime)
		result := big.NewInt(int64(x))
		result.Mul(result, bA)
		return int(result.Mul(result, bA).Add(result, bB).Mod(result, bPrime).Int64() % int64(B))
	}
}

package hashing

type HashingFamily interface {
	GethashingFunction(i int) func(int) int
	GetHashingFunctionNumber() int
}

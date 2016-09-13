package stream

import (
	"math"
	"math/rand"
	"time"
)

func GenerateUniformIntegerStream(N, M int) []int {
	stream := make([]int, N)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < N; i++ {
		stream[i] = r.Intn(M)
	}

	return stream
}

func GenerateExponentialIntegerStream(N, M int) []int {
	stream := make([]int, N)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	lambda := math.MaxFloat64 / float64(M)
	for i := 0; i < N; i++ {
		stream[i] = int(math.Ceil(r.ExpFloat64() / lambda))
	}

	return stream
}

func GenerateNormalIntegerStream(N, M int) []int {
	stream := make([]int, N)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	stdDev := float64(M) / 6.0
	mean := M / 2
	for i := 0; i < N; i++ {
		stream[i] = int(math.Ceil(r.NormFloat64() * stdDev / float64(mean)))
	}

	return stream
}

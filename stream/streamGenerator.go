package stream

import (
	"math/rand"
	"time"
)

func GenerateIntegerStream(N, M int) []int {
	stream := make([]int, N)
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)
	for i := 0; i < N; i++ {
		stream[i] = r.Intn(M)
	}

	return stream
}

// func GenerateStringStream() {

// }

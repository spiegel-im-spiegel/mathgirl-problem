package prob

import (
	"math"
	"math/rand"
	"time"
)

func New(p float64) <-chan bool {
	ch := make(chan bool)
	go func() {
		defer close(ch)
		max := 1000000
		limit := percent(p, max)
		rnd := rand.New(rand.NewSource(time.Now().UnixNano()))
		for {
			n := rnd.Intn(max) + 1
			ch <- n < limit
		}
	}()
	return ch
}

func percent(f float64, max int) int {
	if f < 0 {
		return 0
	}
	if f > 1.0 {
		return max
	}
	return int(math.Floor(f*float64(max) + 0.5))
}

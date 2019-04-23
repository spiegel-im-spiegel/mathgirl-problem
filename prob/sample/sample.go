package main

import (
	"fmt"
	"math"

	"github.com/spiegel-im-spiegel/mathgirl-problem/prob"
)

func probability(try int, ch <-chan bool) int {
	count := 0
	for i := 0; i < try; i++ {
		if <-ch {
			count++
		}
	}
	return count
}

func main() {
	ch := prob.New(0.1)
	min := float64(1)
	max := float64(0)
	sum := float64(0)
	sum2 := float64(0)
	try := 10000
	tryf := float64(try)
	ps := make([]float64, 0, try)
	for i := 0; i < try; i++ {
		count := probability(try, ch)
		p := float64(count) / tryf
		ps = append(ps, p)
		if p < min {
			min = p
		}
		if p > max {
			max = p
		}
		sum += p
		sum2 += p * p
	}
	fmt.Printf("minimum value: %7.5f\n", min)
	fmt.Printf("maximum value: %7.5f\n", max)
	ave := sum / tryf
	fmt.Printf("average: %7.5f\n", ave)
	devi := math.Sqrt(sum2/tryf - ave*ave) //standard deviation
	ct := 0
	for _, p := range ps {
		if ave-devi <= p && p <= ave+devi {
			ct++
		}
	}
	fmt.Printf("standard deviation: %7.5f (%4.1f%%)\n", devi, float64(ct)*100.0/tryf)
}

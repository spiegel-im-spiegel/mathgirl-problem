package main

import (
	"flag"
	"fmt"
	"math"
	"os"
	"strconv"

	"github.com/spiegel-im-spiegel/mathgirl-problem/note257"
)

func probability(ppl *note257.People, tk *note257.TestKit, try int) float64 {
	total := 0
	count := 0
	for i := 0; i < try; i++ {
		psn := ppl.SelectPersion()
		if tk.Inspect(psn) {
			total++
			if psn.Infection() {
				count++
			}
		}
	}
	return float64(count) / float64(total)
}

func main() {
	flag.Parse()
	try, err := strconv.Atoi(flag.Arg(0))
	if err != nil {
		fmt.Fprintln(os.Stderr, err)
		return
	}

	ppl := note257.NewPeople()
	tk := note257.NewTestKit()
	min := float64(1)
	max := float64(0)
	sum := float64(0)
	sum2 := float64(0)
	//try := 1000
	tryf := float64(try)
	ps := make([]float64, 0, try)
	for i := 0; i < try; i++ {
		p := probability(ppl, tk, try*10)
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

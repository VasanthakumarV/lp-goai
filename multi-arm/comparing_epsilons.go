package main

import (
	"fmt"
	"math/rand"
)

// Bandit struct represenets hyp bandit with known return encoded in m,
// the number of times it has run is captured in n, and mean keeps track
// of the average returns from a particular bandit
type Bandit struct {
	n    float64
	m    float64
	mean float64
}

// Pull returns a number that belongs to the gaussian dist of its actual return(m)
func (b *Bandit) Pull() float64 {
	return b.m + rand.NormFloat64()
}

// Update returns the average return for a particular bandit
func (b *Bandit) Update(x float64) {
	b.n += 1.0
	b.mean = (1-(1.0/b.n))*b.mean + ((1.0 / b.n) * x)
}

func main() {
	bandits := []*Bandit{&Bandit{m: 1.0}, &Bandit{m: 2.0}, &Bandit{m: 3.0}}
	// running an experiment for the given list of bandits, the greedy factor for the experiment
	// and the number of iters
	run_experiment(bandits, 0.5, 10)
	for _, b := range bandits {
		fmt.Println(b.mean, b.m)
	}
}

func findMax(bandits []*Bandit) *Bandit {
	max := bandits[0]
	for _, bandit := range bandits {
		if bandit.mean > max.mean {
			max = bandit
		}
	}
	return max
}

func run_experiment(bandits []*Bandit, eps float64, n int) {
	for i := 0; i < n; i++ {
		p := rand.Float64()
		var b *Bandit
		if p < eps {
			b = bandits[rand.Intn(len(bandits))]
		} else {
			b = findMax(bandits)
		}
		x := b.Pull()
		b.Update(x)
	}
}

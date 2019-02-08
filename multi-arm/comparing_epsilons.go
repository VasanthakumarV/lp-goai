package main

import (
	"fmt"
	"math/rand"
)

type Bandit struct {
	n    float64
	m    float64
	mean float64
}

func (b *Bandit) Pull() float64 {
	return b.m + rand.NormFloat64()
}

func (b *Bandit) Update(x float64) {
	b.n += 1.0
	b.mean = (1-(1.0/b.n))*b.mean + ((1.0 / b.n) * x)
}

func findMax(bandits []*Bandit) *Bandit {
	max := bandits[rand.Intn(len(bandits))]
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

func main() {
	bandits := []*Bandit{&Bandit{m: 1.0}, &Bandit{m: 2.0}, &Bandit{m: 3.0}}
	run_experiment(bandits, 0.5, 10)
	for _, b := range bandits {
		fmt.Println(b.mean, b.m)
	}
}

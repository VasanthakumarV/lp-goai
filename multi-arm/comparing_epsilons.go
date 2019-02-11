package main

import (
	"fmt"
	"log"
	"math/rand"
	"os"

	"gonum.org/v1/plot"
	"gonum.org/v1/plot/plotter"
)

// Bandit struct represenets hyp bandit with known return(reward) encoded in m,
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
	runExperiment(bandits, 0.5, 1000)
	for _, b := range bandits {
		fmt.Println(b.mean, b.m)
	}
}

func runExperiment(bandits []*Bandit, eps float64, n int) {
	data := make(plotter.XYs, n)
	for i := 0; i < n; i++ {
		p := rand.Float64()
		var b *Bandit
		if p < eps {
			b = bandits[rand.Intn(len(bandits))]
		} else {
			b = findMax(bandits)
		}
		x := b.Pull()

		data[i].X = float64(i)
		data[i].Y = x
		b.Update(x)
	}

	err := plotData("out.png", data)
	if err != nil {
		log.Fatalf("error creating chart: %v", err)
	}
}

func plotData(path string, data plotter.XYs) error {
	f, err := os.Create(path)
	if err != nil {
		return fmt.Errorf("could not create png file: %v", err)
	}
	defer f.Close()

	p, err := plot.New()
	if err != nil {
		return fmt.Errorf("could not create plot: %v", err)
	}

	s, err := plotter.NewLine(plotter.XYs(data))
	if err != nil {
		return fmt.Errorf("could not create scatter: %v", err)
	}
	p.Add(s)

	wt, err := p.WriterTo(512, 512, "png")
	if err != nil {
		return fmt.Errorf("could not create wtiter: %v", err)
	}

	_, err = wt.WriteTo(f)
	if err != nil {
		return fmt.Errorf("could not write to png file: %v", err)
	}

	return nil

}

// findMax returns the bandit with the highest mean for rewards
func findMax(bandits []*Bandit) *Bandit {
	max := bandits[0]
	for _, bandit := range bandits {
		if bandit.mean > max.mean {
			max = bandit
		}
	}
	return max
}

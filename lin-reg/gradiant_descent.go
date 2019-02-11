package main

import (
	"gonum.org/v1/gonum"
)

type Opt struct {
	mInit, cInit, learningRate float64
	iters                      int
}

func gradient(data mat.NewDense, options Opt) {
	mCurr := options.mInit
	cCurr := options.cInit
	for i := 0; i < options.iters; i++ {
		yCurrent := (mCurr )
	}
}

func main() {
	data := mat.NewDense(2, 2, []float64{1, 1, 2, 2})
	options := Opt{learningRate: 0.01, iters: 1000}

	gradient(, options)

}

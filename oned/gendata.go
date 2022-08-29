package oned

import (
	"math/rand"
)

func gendata(ndata int, mean, sd float64) []float64 {
	var data []float64
	for i := 0; i < ndata; i++ {
		data = append(data, sd*rand.NormFloat64()+mean)
	}

	return data
}

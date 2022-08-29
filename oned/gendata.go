package oned

import (
	"log"
	"math"
	"math/rand"
)

const m_exponent = 2.0

func GenData(ndata int, mean, sd float64) []float64 {
	var data []float64
	for i := 0; i < ndata; i++ {
		data = append(data, sd*rand.NormFloat64()+mean)
	}

	return data
}

func UpdateCenter(data, weights []float64) float64 {
	tw := 0.0
	c := 0.0

	for k, w := range weights {
		tw += w
		c += data[k] * w
	}

	return c / tw
}

// conventional fuzzy c-means
func UpdateMemFuncF(data []float64, c0, c1 float64) ([]float64, []float64, []float64, []float64) {
	var u0, u1, dist0, dist1 []float64

	for _, d := range data {
		d0 := (d - c0) * (d - c0)
		d1 := (d - c1) * (d - c1)
		dist0 = append(dist0, d0)
		dist1 = append(dist1, d1)
		u0 = append(u0, 1.0/(1.0+math.Pow(d0/d1, 1/(m_exponent-1))))
		u1 = append(u1, 1.0/(1.0+math.Pow(d1/d0, 1/(m_exponent-1))))
	}

	return dist0, dist1, u0, u1
}

// Fuzzy-Crisp
func UpdateMemFuncFC(data []float64, c0, c1 float64) ([]float64, []float64, []float64, []float64) {
	var u0, u1, dist0, dist1 []float64
	a0 := 0.5
	a1 := 0.5

	var tu0, tu1 float64

	for _, d := range data {
		d0 := (d - c0) * (d - c0)
		d1 := (d - c1) * (d - c1)
		dc := (1.0 + a0 + a1) / (a0/d0 + a1/d1)

		if d0 < dc && d1 < dc {
			tu0 = a0 * (dc/d0 - 1)
			tu1 = a1 * (dc/d1 - 1)
		} else if d0 < dc {
			tu0 = 1.0
			tu1 = 0.0
		} else if d1 < dc {
			tu0 = 0.0
			tu1 = 1.0
		} else {
			log.Fatal("UpdateMemFuncFC: Cannot happen")
		}
		dist0 = append(dist0, d0)
		dist1 = append(dist1, d1)
		u0 = append(u0, tu0)
		u1 = append(u1, tu1)
	}

	return dist0, dist1, u0, u1
}

// weight for updating centers (conventional FCM)
func MF2WeightF(memfunc []float64) []float64 {
	var w []float64
	for _, u := range memfunc {
		w = append(w, math.Pow(u, m_exponent))
	}
	return w
}

// weight for updating centers (Fuzzy-Crisp clustering)
func MF2WeightFC(memfunc []float64) []float64 {
	var w []float64
	for _, u := range memfunc {
		v := 0.5*u*u + u
		w = append(w, v)
	}
	return w
}

func GetTargetFunc(data, dist0, dist1, w0, w1 []float64) float64 {
	target := 0.0
	for k := range data {
		target += w0[k]*dist0[k] + w1[k]*dist1[k]
	}

	return target
}

func StepFCM(data []float64, c0, c1 float64) (float64, float64, float64, []float64, []float64) {
	dist0, dist1, mf0, mf1 := UpdateMemFuncF(data, c0, c1)
	w0 := MF2WeightF(mf0)
	w1 := MF2WeightF(mf1)

	target := GetTargetFunc(data, dist0, dist1, w0, w1)

	c0 = UpdateCenter(data, w0)
	c1 = UpdateCenter(data, w1)

	return target, c0, c1, mf0, mf1
}

func StepFCCM(data []float64, c0, c1 float64) (float64, float64, float64, []float64, []float64) {
	dist0, dist1, mf0, mf1 := UpdateMemFuncFC(data, c0, c1)
	w0 := MF2WeightFC(mf0)
	w1 := MF2WeightFC(mf1)

	target := GetTargetFunc(data, dist0, dist1, w0, w1)

	c0 = UpdateCenter(data, w0)
	c1 = UpdateCenter(data, w1)

	return target, c0, c1, mf0, mf1
}

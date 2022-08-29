package main

import (
	"fmt"
	"log"

	"github.com/arkinjo/fccm/oned"
)

func cluster(flag int, data []float64, c0, c1 float64) {
	var u0, u1 []float64
	var tf, nc0, nc1 float64
	for n := 0; n <= 20; n++ {
		if flag == 0 {
			tf, nc0, nc1, u0, u1 = oned.StepFCM(data, c0, c1)
		} else {
			tf, nc0, nc1, u0, u1 = oned.StepFCCM(data, c0, c1)
		}
		fmt.Printf("Step%d\t%d\t%f\t%f\t%f\n", flag, n, tf, c0, c1)
		c0 = nc0
		c1 = nc1

	}
	for k, d := range data {
		fmt.Printf("Data%d\t%d\t%f\t%f\t%f\n", flag, k, d, u0[k], u1[k])
	}

}

func main() {
	c0 := -5.0
	c1 := 5.0

	data1 := oned.GenData(1000000, c0, 1.0)
	data2 := oned.GenData(1000, c1, 1.0)
	data := append(data1, data2...)

	cluster(0, data, c0, c1)
	cluster(1, data, c0, c1)

	log.Println("clusters:", c0, c1)
}

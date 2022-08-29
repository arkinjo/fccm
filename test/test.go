package main

import (
	"fmt"
	"github.com/arkinjo/fccm/oned"
)

func main() {
	data := oned.gendata(100, -1, 1)
	for i, d := range data {
		fmt.Println(i, d)
	}
}

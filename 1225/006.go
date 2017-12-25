package main

import (
	"fmt"
)

func compare(x,y float64) float64 {
	if x < y {
		return x

	}
	return y
}

func main() {
	fmt.Println(compare(12,5))
}


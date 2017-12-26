package main

import (
	"fmt"
	"math"
)

func pow2(x,n,lim float64) float64  {
	if v:=math.Pow(x,n); v<lim{
		fmt.Printf("%g <= %g \n",v,lim)
		return v
	} else {
		fmt.Printf("%g >= %g \n",v,lim)
	}
	return lim
}

func main() {
	fmt.Println(
		pow2(2,3,10),
	)
	fmt.Println(
		pow2(3,3,20),
	)
}
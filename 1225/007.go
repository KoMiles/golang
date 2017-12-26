package main

import "fmt"
import "math"

func Sqrt(x float64) float64 {
	const E = 0.00001
	z := float64(1)
	k := float64(0)
	for ; ; z = z - ((z*z -x ) / (2*z)) {
		if z-k <=E && z-k <=-E{
			return z
		}
		k = z
	}
}

func main()  {
	fmt.Println(Sqrt(5))
	fmt.Println(math.Sqrt(5))
}

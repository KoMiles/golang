package main

import "fmt"

func split(sum int)(x, y int) {
	x = sum * 4 /8
	y = sum - x
	return
}

func main() {
	fmt.Println(split(10))
}
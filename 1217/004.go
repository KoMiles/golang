package main

import "fmt"

func add(x int, y int)  int{
	return x + y
}

func addOne(x,y int) int{
	return x+y
}

func main() {
	fmt.Println(addOne(8,10))
}
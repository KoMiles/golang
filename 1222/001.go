package main

import "fmt"

const PI = 3.14

func main() {
	const World = "世界"
	var name = "孔明"
	fmt.Println("Hello", World)
	fmt.Println("Happy",PI, "Day", name)

	const Truth = true
	fmt.Println("Go rules ?", Truth)
}

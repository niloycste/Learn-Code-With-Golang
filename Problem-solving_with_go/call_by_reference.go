package main

import "fmt"

func increment(x *int) {
	*x++
}

func main() {
	x := 5
	increment(&x)
	fmt.Println("x after increment:", x) // x will be 6
}

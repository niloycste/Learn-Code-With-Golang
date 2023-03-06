package main

import "fmt"

func add(a int, b int) int {
	return a + b
}

func main() {
	result := add(3, 4)
	fmt.Println("The sum of 3 and 4 is:", result)
}

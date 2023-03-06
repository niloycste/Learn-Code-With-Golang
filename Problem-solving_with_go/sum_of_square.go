package main

import "fmt"

func main() {
	var n int
	var sum int

	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&n)

	i := 1
	for i <= n {
		sum = sum + (i * i)
		i++
	}

	fmt.Println(sum)
}

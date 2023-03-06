package main

import "fmt"

func main() {
	fmt.Println("Enter the size of an array:")
	var size int
	fmt.Scan(&size)
	fmt.Println("Enter values for an array:")
	var a []int
	for i := 0; i < size; i++ {
		var input int
		fmt.Scan(&input)
		a = append(a, input)
	}
	var min = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] < min {
			min = a[i]
		}
	}
	fmt.Println("Maximum value:", min)
}

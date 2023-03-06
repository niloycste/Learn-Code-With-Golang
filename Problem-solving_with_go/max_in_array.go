package main

import "fmt"

func main() {
	fmt.Println("max and minimum value from a array")
	var a = []int{4, 3, 5, 2, 1, 6}
	var max = a[0]
	for i := 1; i < len(a); i++ {
		if a[i] > max {
			max = a[i]
		}
	}
	fmt.Println(max)

}

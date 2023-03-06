package main

import "fmt"

func main() {
	a := []int{1, 2, 3, 4, 5}
	fmt.Println("Original array:", a)

	// insert element 6 at index 2
	n := len(a) - 1
	for i := n; i > 2; i-- {
		a[i] = a[i-1]
	}
	a[2] = 6
	fmt.Println("Array after inserting 6 at index 2:", a)
}

/*
func main() {
	a := []int{1, 2, 3}
	fmt.Println("Original array:", a)

	// insert element 4 at index 1
	a = append(a[:1], append([]int{4}, a[1:]...)...)
	fmt.Println("Array after inserting 4 at index 1:", a)
}
*/

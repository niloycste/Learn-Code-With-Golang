package main

import "fmt"

func main() {
	fmt.Println("enter the size of an array")
	var size int
	fmt.Scan(&size)
	fmt.Println("enter the values of an array")
	var a []int
	for i := 0; i < size; i++ {
		var input int
		fmt.Scan(&input)
		a = append(a, input)
	}

	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if (a[i] == a[j]) && (i != j) {
				fmt.Println("Duplicates elemenets are", a[j])
			}
		}
	}

}

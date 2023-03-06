package main

import "fmt"

func main() {
	fmt.Println("find out the second largest value from a given array")
	fmt.Println("Enter a Size of a array")
	var size int
	fmt.Scan(&size)
	fmt.Println("enter the value of a array")
	var a []int
	for i := 0; i < size; i++ {
		var input int
		fmt.Scan(&input)
		a = append(a, input)
	}

	var temp int
	for i := 0; i < len(a); i++ {
		for j := i + 1; j < len(a); j++ {
			if a[i] < a[j] {
				temp = a[i]
				a[i] = a[j]
				a[j] = temp
			}
		}
	}
	fmt.Println("second largest value is", a[1])

}

//2nd smallest number is
// "if a[i] > a[j]"

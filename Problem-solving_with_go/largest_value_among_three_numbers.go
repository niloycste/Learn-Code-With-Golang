package main

import "fmt"

func main() {
	var num1, num2, num3 int
	fmt.Println("Enter three numbers:")
	fmt.Scan(&num1, &num2, &num3)

	var largest int
	if num1 > num2 && num1 > num3 {
		largest = num1
	} else if num2 > num1 && num2 > num3 {
		largest = num2
	} else {
		largest = num3
	}

	fmt.Println("The largest value is:", largest)
}

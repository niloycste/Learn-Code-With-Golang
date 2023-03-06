package main

import "fmt"

func main() {
	var num1, num2 int
	fmt.Println("Enter the Two numbers")
	fmt.Scan(&num1, &num2)
	x := num1
	y := num2
	for y != 0 {
		rem := x % y
		x = y
		y = rem

	}

	fmt.Println("The GCD of 12 and 18 is:", x)
	lcm := (num1 * num2) / x
	fmt.Println("The lcm of 12 and 18 is", lcm)
}

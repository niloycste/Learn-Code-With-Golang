package main

import "fmt"

func main() {
	var num int
	fmt.Println("Enter a number:")
	fmt.Scan(&num)

	var firstDigit int
	for firstDigit = num; firstDigit > 9; firstDigit = firstDigit / 10 {
	}
	lastDigit := num % 10

	fmt.Println("The first digit of", num, "is", firstDigit)
	fmt.Println("The last digit of", num, "is", lastDigit)
}

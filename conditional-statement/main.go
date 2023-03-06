package main

import "fmt"

func main() {
	fmt.Println("welcome to conditional statement")

	number := 0
	if number > 0 {
		fmt.Println("positive")
	} else if number < 0 {
		fmt.Println("negative")
	} else {
		fmt.Println("zero")
	}
	//even-odd program
	var value int
	fmt.Printf("Enter any value:")
	fmt.Scan(&value)

	if value%2 == 0 {
		fmt.Println("even")
	} else {
		fmt.Println("odd")
	}
}

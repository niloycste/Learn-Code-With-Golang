package main

import "fmt"

func main() {
	fmt.Println("welcome to golang switch")

	var classnumber int
	fmt.Printf("Enter your class:\n")
	fmt.Scan(&classnumber)

	switch classnumber {
	case 1, 2, 3, 4, 5:
		fmt.Println("primary")
	case 6, 7, 8, 9, 10:
		fmt.Println("secondary")
	case 11, 12:
		fmt.Println("higher secondary")
	default:
		fmt.Println("university")
	}
}

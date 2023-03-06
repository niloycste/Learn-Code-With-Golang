package main

import "fmt"

func main() {
	fmt.Println("welcome to golang switch")

	var digit int
	fmt.Printf("Enter a digit from 0-9")
	fmt.Scan(&digit)

	switch digit {
	case 0:
		fmt.Printf("zero\n")
	case 1:
		fmt.Printf("one\n")
	case 2:
		fmt.Printf("two\n")
	case 3:
		fmt.Printf("three\n")
	case 4:
		fmt.Printf("four\n")
	case 5:
		fmt.Printf("five\n")
	case 6:
		fmt.Printf("six\n")
	case 7:
		fmt.Println("seven")
	default:
		fmt.Println("not a valid reason")
	}
}

//thre is no continue or break statement on switch in golang.
//if we use string then we have to use "" in case like case "0", case "1 this way

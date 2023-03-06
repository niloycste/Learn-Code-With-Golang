package main

import "fmt"

func main() {
	fmt.Println("swap two numbers using 3rd vairable")

	x := 10
	y := 20
	x = x + y
	y = x - y
	x = x - y
	fmt.Println(x)
	fmt.Println(y)

}

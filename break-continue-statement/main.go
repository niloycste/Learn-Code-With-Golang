package main

import "fmt"

func main() {
	fmt.Println("welcome to golang break and continue")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			continue //if a number a divisible by 2 that will not print because of "continue"
		}
		fmt.Println(i)
	}
}

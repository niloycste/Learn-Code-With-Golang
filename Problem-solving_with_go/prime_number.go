package main

import "fmt"

func main() {
	fmt.Println("Prime numbers between 1 and 20:")
	var temp = 0
	var n int
	fmt.Println("Total values are")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		for j := 2; j <= i-1; j++ {
			if i%j == 0 {
				temp++

			}
		}
		if temp == 0 {
			fmt.Println("prime numbers are", i)
		} else {
			temp = 0
		}
	}
}

package main

import "fmt"

func main() {
	var x = 10
	var y = 20
	var sum int
	ptr1 := &x
	ptr2 := &y
	sum = *ptr1 + *ptr2
	fmt.Println("sum of two pointer", sum)
}

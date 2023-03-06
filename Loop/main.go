package main

import "fmt"

func main() {
	fmt.Println("loop in golang")

	for x := 1; x < 10; x++ {
		if x%2 == 0 {

			fmt.Println(x)
		}
	}

}

// we can use for loop as a while loop
func f() {
	var y int
	y = 10

	for y <= 100 {
		fmt.Println(y)
		y = y + 2
	}

}

//in golang there is only for loop, while, do while doesn;t exist.
//for range loop, we will use it for array and slicing

func ni() {
	s := []int{10, 55, 88, 777, 69, 96}

	for index, value := range s {
		fmt.Println(index, value)
	}
}

//if we dont want to print value not index then

func nil() {
	s := []int{10, 55, 88, 777, 69, 96}

	for _, value := range s {
		fmt.Println(value)
	}
}

//if we want to print index but not value then

func niloy() {
	s := []int{10, 55, 88, 777, 69, 96}

	for index, _ := range s {
		fmt.Println(index)
	}
}

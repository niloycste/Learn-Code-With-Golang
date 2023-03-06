package main

import "fmt"

func main() {
	fmt.Println("welcome to slices and its length and capacity ")
	y := make([]int, 0, 2) //len=0, capacity=2
	fmt.Println(y)
	fmt.Printf("%T", y)
	fmt.Println("length", len(y))
	fmt.Println("capacity", cap(y))

	y = append(y, 17)
	fmt.Println(y)
	fmt.Printf("%T\n", y)
	fmt.Println("length", len(y))
	fmt.Println("capacity", cap(y))

	y = append(y, 19)
	fmt.Println(y)
	fmt.Printf("%T", y)
	fmt.Println("length", len(y))
	fmt.Println("capacity", cap(y))

	y = append(y, 21) //fill the previous capacity of 2 , now it assain another caapcity 4 to assain a value
	fmt.Println(y)
	fmt.Printf("%T", y)
	fmt.Println("length", len(y))
	fmt.Println("capacity", cap(y))

}

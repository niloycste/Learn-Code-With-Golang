package main

import "fmt"

func main() {
	fmt.Println("welcome to  pointers")

	//var ptr *int
	//fmt.Println("value of pointer is", ptr)
	var mynumber = 25
	var ptr = &mynumber
	fmt.Println("value of actual pointer is:", ptr)       //address
	fmt.Println("value of actual pointer is:", &mynumber) //address

	fmt.Println("value of actual pointer is:", *ptr) //actual value

	*ptr = *ptr * 2
	fmt.Println("new value is:", mynumber)

}

package main

import "fmt"

func main() {
	fmt.Println("welcome to structs in golang")

	//structs and class is same. and in golang there is no inheritance,super or parent
	details := User{"Niloy", "niloy@yahoo.com", true, 27}
	fmt.Println("niloy's details is:", details)
	fmt.Printf("niloy's details is: %+v\n", details)
}

type User struct {
	Name   string
	Emain  string
	Status bool
	Age    int
}

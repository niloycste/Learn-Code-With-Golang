package main

import "fmt"

func main() {
	fmt.Println("golang structs")

	niloy := Student{1401068, "comilla", 28} //object creation
	piyas := Student{1401017, "tangail", 28}
	nayan := Student{1401040, "savar", 28}
	fmt.Printf("details are:%+v\n", niloy)
	fmt.Printf("details are:%+v\n", nayan)
	fmt.Printf("details are:%+v\n", piyas)
	fmt.Printf("details are:%+v\n", niloy.id)
	fmt.Printf("details are:%+v\n", piyas.address)

}

type Student struct {
	id      int
	address string
	age     int
}

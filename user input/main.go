package main

import "fmt"

func main() {
	var studentname [5]string
	var i int
	for i = 0; i < len(studentname); i++ {
		fmt.Println("enter a student name")
		fmt.Scan(&studentname[i])

	}
	fmt.Println(studentname)

}

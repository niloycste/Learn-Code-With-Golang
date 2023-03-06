package main

import "fmt"

func main() {
	fmt.Println("remove value from slices based on index")
	courses := []string{"C", "C++", "JAVA", "RUST", "GOLANG", "PYTHON"}
	fmt.Println(courses)
	index := 2 // we want to remove index 2(JAVA)
	courses = append(courses[:index], courses[index+1:]...)
	fmt.Println("values without index num:2", courses)

}

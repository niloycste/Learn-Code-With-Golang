package main

import "fmt"

func displayMessage(countryName string) {
	//countryname-variable name and then we have to use datatype
	fmt.Println("print country name:", countryName)

}

// single paramatre
func square(num int) {
	result := num * num
	fmt.Println(result)
}

// multiple parametre
func add(num1 int, num2 int) {
	fmt.Println(num1 + num2)
}
func sub(num1 int, num2 int) int {
	//we use int after the argument because wehn we use "return" then we have to assiagn its datatype. thats why we use "int after sub(num1 int, num2 int) int "
	result := num1 - num2
	return result // we use return wehen we dont want to print
}

func message() string {
	return "hello niloy"
}

func main() {
	displayMessage("Bangladesh") // call the function
	displayMessage("America")
	displayMessage("India")
	square(10)
	add(5, 10)
	fmt.Println(sub(50, 25))
	text := message() //we can assian function to another variable
	fmt.Println(text)

}

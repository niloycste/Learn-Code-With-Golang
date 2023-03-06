package main

import "fmt"

func main() {
	var sentence string
	fmt.Println("Enter a sentence:")
	fmt.Scan(&sentence)

	var vowelCount, consonantCount int
	for _, char := range sentence {
		/*This is a for loop in Go (also known as Golang) that iterates through
		each character in a string called "sentence". The variable "char" is
		assigned each character in the string during each iteration.
		 The first variable, "_", is a blank identifier and is used when the index
		 of the character in the string is not needed.*/
		switch {
		case char == 'a' || char == 'e' || char == 'i' || char == 'o' || char == 'u' || char == 'A' || char == 'E' || char == 'I' || char == 'O' || char == 'U':
			vowelCount++
		case char >= 'a' && char <= 'z' || char >= 'A' && char <= 'Z':
			consonantCount++
		}
	}
	fmt.Println("Number of vowels in the sentence: ", vowelCount)
	fmt.Println("Number of consonants in the sentence: ", consonantCount)
}

//why we use (_) here
/*
The variable "_" is a blank identifier, and it is used when the index of the character
in the string is not needed. In this case, the program only needs to access the value
of each character in the string, not its position, so the index is not used and the blank identifier
is used instead. It's a way of saying that the value of the index is not important to use,
and it will not be used in the loop, it's more like a placeholder,
it's not required to use it but it's a good practice when the index is not needed.
*/

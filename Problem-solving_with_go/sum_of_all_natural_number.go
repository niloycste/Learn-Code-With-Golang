package main

import "fmt"

func main() {
	var n int
	var sum = 0

	fmt.Print("Enter a positive integer: ")
	fmt.Scan(&n)

	for i := 1; i <= n; i++ {
		sum += i
	}

	fmt.Println("Sum of first", n, "natural numbers is:", sum)
}

/*package main

import "fmt"

func main() {
    var n int
    fmt.Print("Enter a positive integer: ")
    fmt.Scan(&n)
    sum:=n*(n+1)/2
    fmt.Println("Sum of first", n, "natural numbers is:", sum)
}
*/
// using another for loop(while loop)
/*
package main

import "fmt"

func main() {
    var n int
    var sum int

    fmt.Print("Enter a positive integer: ")
    fmt.Scan(&n)

    i := 1
    for i <= n {
        sum += i
        i++
    }

    fmt.Println("Sum of first", n, "natural numbers is:", sum)
}
*/

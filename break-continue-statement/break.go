package main

import "fmt"

func main() {
	fmt.Println("welcome to golang break and continue")

	for i := 1; i <= 10; i++ {
		if i%2 == 0 {
			break

			/*when some value divisble by 2 then it come out from loop.
			like 2 is divisible by 2 when it found 2 it come out from
			the loop and it will not enter into the loop
			*/

		}
		fmt.Println(i)
	}
}

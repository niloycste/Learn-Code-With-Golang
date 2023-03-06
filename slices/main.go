package main

import (
	"fmt"
	"sort"
)

func main() {
	fmt.Println("welcome to slices on golang")

	fruitlist := []string{"apple", "banana", "mango", "peach"}
	fmt.Printf("type of fruitlist is:%T\n", fruitlist)
	fmt.Println("type of fruitlist is:", fruitlist)

	fruitlist = append(fruitlist, "watermelon", "tomato")
	fruitlist = append(fruitlist[1:]) //slices on existing array or slice
	fruitlist = append(fruitlist[:3]) //not inclusive so index 3 will not be printed
	fruitlist = append(fruitlist[2:4])

	fmt.Println("type of fruitlist is:", fruitlist)

	//make keyword

	highscores := make([]int, 4)
	highscores[0] = 231
	highscores[1] = 234
	highscores[2] = 237
	highscores[3] = 239
	//highscores[4] = 250 //will not print cause out of index array
	fmt.Println("int value is:", highscores)

	highscores = append(highscores, 432, 244, 31, 212) //this will work because of append
	fmt.Println("int value is:", highscores)
	fmt.Println(sort.IntsAreSorted(highscores)) //give boolean value

	sort.Ints(highscores) //sort value on ascending order
	fmt.Println(sort.IntsAreSorted(highscores))
	fmt.Println(highscores)
}

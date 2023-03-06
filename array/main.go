package main

import "fmt"

func main() {
	fmt.Println("welcome to array in golang")

	var fruitlist [4]string
	fruitlist[0] = "jackfruit"
	fruitlist[1] = "mango"
	fruitlist[2] = "watermelon"
	fmt.Println("fruitlist is:", fruitlist)

	fmt.Println("fruitlist is:", len(fruitlist))

	var veglist = [5]string{"potato", "beans", "mushroom", "cornflower"}
	fmt.Println("vegetable list is:", veglist)
}

package main

import "fmt"

type bot struct {
	lang string
	id   int
}

type test struct {
}

func (b bot) speak() {

	fmt.Println("I speak", b.lang, "and my id is", b.id)
}

func main() {

	mechbot := bot{"Bengali", 300}
	chatbot := bot{"English", 500}
	//t:=test{}

	mechbot.speak()
	chatbot.speak()
}

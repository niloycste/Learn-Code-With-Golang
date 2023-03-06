package main

import "fmt"

func main() {
	fmt.Println("welcome to golang Maps")

	languages := make(map[string]string) //need to use map keyword into make keyword

	languages["js"] = "javascript"
	languages["py"] = "python"
	languages["go"] = "golang"
	languages["rs"] = "rust"
	fmt.Println("list of all languages:", languages)
	fmt.Println("Py shorts for:", languages["py"])
	delete(languages, "rs") //delete a value of rust
	fmt.Println("list of all languages:", languages)

}

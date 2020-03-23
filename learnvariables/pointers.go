package main

import "fmt"

func main() {

	var firstPtr *string = new(string)
	*firstPtr = "Yoda"
	fmt.Println(*firstPtr)

	secondName := "Anakin"
	fmt.Println(secondName)
	ptr := &secondName
	fmt.Println(ptr, *ptr)
	secondName = "Darth Vader"
	fmt.Println(ptr, *ptr)
}

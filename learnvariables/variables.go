package main

import "fmt"

var android, ios, golang bool

func main() {
	var x float32 = 5.15
	fmt.Println(android, ios, golang, x)

	firstname := "Yoda"
	fmt.Println(firstname)

	c := complex(3, 4)
	fmt.Println(c)

	r, i := real(c), imag(c)
	fmt.Println(r, i)
}

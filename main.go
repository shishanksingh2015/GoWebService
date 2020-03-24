package main

import (
	"fmt"

	"github.com/shishanksingh2015/GoWebService/models"
)

func main() {
	u := models.User{
		ID:        2,
		FirstName: "Darth",
		LastName:  "Vader",
	}
	fmt.Println(u)
}

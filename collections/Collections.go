package main

import "fmt"

func main() {

	var arr [3]int
	arr[0] = 1
	arr[1] = 5
	arr[2] = 2
	fmt.Print(arr)

	newArr := [3]int{5, 6, 7}
	fmt.Println(newArr)

	var slice = arr[:]
	fmt.Println(slice)
	var slice1 = []int{1, 2, 3}
	fmt.Println(slice1)
	slice1 = append(slice1, 45, 65, 123, 13, 12, 12)
	fmt.Println(slice1)
	s2 := slice1[1:]
	s3 := slice1[:2]
	s4 := slice1[1:2]
	fmt.Println(s2, s3, s4)
}

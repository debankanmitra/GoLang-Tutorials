package main

import "fmt"

func main() {
	var ptr *int
	fmt.Println("value of pointer is ", ptr)

	num := 23

	var pointer = &num
	fmt.Println("referance of the num or address of the variable in memory", pointer)
	fmt.Println("pointer with * ", *pointer)

	*pointer = *pointer * 2
	fmt.Println("new value is ", *pointer)
}

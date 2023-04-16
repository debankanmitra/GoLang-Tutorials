package main

import "fmt"

func main() {
	// there is no inheritance in golang

	debankan := User{"Dev", "debankanmitra@gmail.com", true, 21}
	fmt.Printf("User deatils is %+v\n", debankan)
}

type User struct { // this struct is similar to c
	Name   string
	Email  string
	Status bool
	Age    int
}

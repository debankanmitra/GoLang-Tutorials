package main

import "fmt"

func main() {
	defer fmt.Println("hi my name is debankan") // statement with defer keyword will be executed last
	fmt.Println("i lives in kolkata")

	defer fmt.Println("one")
	defer fmt.Println("two")
	defer fmt.Println("three") // execution will be lst in firt out

}

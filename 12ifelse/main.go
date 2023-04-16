package main

import "fmt"

func main() {
	logincount := 12
	var result string

	if logincount < 10 {
		result = "not regular"
	} else if logincount == 10 {
		result = "medium"
	} else {
		result = " regular"
	}

	fmt.Println("the result is ", result)

	// direct

	if num := 5; num < 10 {
		num += 5
		fmt.Println("the num is ", num)

	} else {
		num -= 5
		fmt.Println("the num is ", num)

	}
}

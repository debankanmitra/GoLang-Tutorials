package main

import "fmt"

func main() {
	addtwo(5, 8)

	// func foo(){               /*function inside function is not allowed in golang*/
	// 	fmt.Println("hi all")
	// }

	fmt.Println("hi all")

	result := multiply(5, 4)
	fmt.Println(result)

	todl, message := proadder(2, 3, 6, 8)
	fmt.Println(message, ":", todl)
}

func addtwo(a int, b int) {
	fmt.Println(a + b)
}

func multiply(a int, b int) int { // outside int defines what type of result to return these are called function signatures
	return a * b
}

func proadder(values ...int) (int, string) { // we loop through because values are now working as slices and it is like spread operator in js
	total := 0

	for _, val := range values {
		total += val
	}

	return total, "hi this is total"
}

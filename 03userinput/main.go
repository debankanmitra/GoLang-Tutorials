package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	welcome := "hii welcome"
	fmt.Println(welcome)

	fmt.Println("Enter anything you like")
	reader := bufio.NewReader(os.Stdin)

	// comma ok || error error
	input, _ := reader.ReadString('\n') // _ it means we don't care about the error
	fmt.Println("entered thing is ", input)
}

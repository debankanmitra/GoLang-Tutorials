package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)

func main() {
	fmt.Println("enter a number")
	readinput := bufio.NewReader(os.Stdin) // referance of a reader
	input, _ := readinput.ReadString('\n') // how the reader will read
	fmt.Println("number is ", input)

	num, err := strconv.ParseFloat(strings.TrimSpace(input), 32) // comma ok syntax
	fmt.Println(err)                                             // strconv.ParseFloat: parsing "90\n": invalid syntax
	fmt.Println("input after adding 1:", num+1)

}

package main

import (
	"fmt"
)

// numbers := 700, this voldoras operator doesn't run outside function so we need to use var keyword then

const Logintoken string = "SFDOR8F SGFDG 99FG" // Mentioning a varibale with capital letter makes it publicly accessible

func main() {
	fmt.Println("variables")

	var username string = "debankan"
	fmt.Println(username)
	fmt.Printf("Type :%T \n", username) // don't print in new line so we gave \n

	var isLoggedIn bool = true
	fmt.Println(isLoggedIn)
	fmt.Printf("Type: %T \n", isLoggedIn)

	var smallval uint8 = 255 // set of all unsigned 8-bit integers only can holds upto (0-255) , int8 is signed 8-bit numbers , we use uint when we do something very specific to os
	fmt.Println(smallval)
	fmt.Printf("Type: %T \n", smallval)

	var smallFloat float32 = 255.9898349530403543905634
	fmt.Println(smallFloat)
	fmt.Printf("Type :%T \n", smallFloat)

	// default values
	var deval int
	fmt.Println(deval)
	fmt.Printf("Type: %T \n", deval)

	// implicit type
	var website = "google.com" // lexer takes car for its types
	fmt.Println(website)
	fmt.Printf("Type: %T \n", website)

	// no var style
	numbers := 700
	fmt.Println(numbers)
	fmt.Printf("Type: %T \n", numbers)

	// publicly accessible
	fmt.Println(Logintoken)

}

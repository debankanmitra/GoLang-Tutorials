package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	dicenum := rand.Intn(6) + 1
	fmt.Println("value of dice number", dicenum)

	switch dicenum {
	case 1:
		fmt.Println("dev ")
	case 2:
		fmt.Println("ravi ")
	case 3:
		fmt.Println("smith")
		fallthrough
	case 4:
		fmt.Println("rahul ")
		fallthrough
	case 5:
		fmt.Println("john ")
	case 6:
		fmt.Println("kimil ")
	}
}

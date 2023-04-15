package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
	//"math/rand"
)

func main() {
	// golang treats random number into two separate categories
	//       1) teh random number generated by math ( not so efficeint )
	//       2) gemetric random number by crptography algorithm

	// math module
	//rand.Seed(time.Now().UnixNano())
	//fmt.Println(rand.Intn(10)) // it will always give same number unless we provide a seeed which keeps changing like time

	// random number from crypto
	randomnum, _ := rand.Int(rand.Reader, big.NewInt(10))
	fmt.Println(randomnum)

}

package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"os"
)

func main() {
	content := "writing text to files"

	// creating a file
	file, err := os.Create("./file.txt")
	if err != nil {
		panic(err) // panic shuts down execution of program
	}

	// writing file
	length, err := io.WriteString(file, content)
	// if err != nil {
	// 	panic(err)
	// }
	checkNilError(err) // to avoid repetation we using function
	fmt.Println("length is ", length)
	defer file.Close()

	readfile("./file.txt")
}

func readfile(filename string) {
	databytes, err := ioutil.ReadFile(filename) // data is read in byte format
	checkNilError(err)                          // to avoid repetation we using function

	fmt.Println("Data is ", string(databytes)) // byte to string
}

func checkNilError(err error) {
	if err != nil {
		panic(err)
	}
}

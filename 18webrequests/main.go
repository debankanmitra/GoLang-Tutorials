package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

const url string = "https://lco.dev/"

func main() {
	response, err := http.Get(url)

	if err != nil {
		panic(err)
	}

	fmt.Printf("response type is: %T\n", response)
	defer response.Body.Close() // caller's responsibitlity to close get request connection

	fmt.Println("response is ", response.Header)

	data, _ := ioutil.ReadAll(response.Body)
	fmt.Println(string(data))
}

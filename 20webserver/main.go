package main

import (
	"fmt"
	"net/http"
)

func main() {
	const myurl = "https://lco.dev/"

	// GET request
	response, err := http.Get(myurl)
	if err != nil {
		panic(err)
	}
	defer response.Body.Close()
	fmt.Println(response.StatusCode)

}

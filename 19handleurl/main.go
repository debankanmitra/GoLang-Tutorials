package main

import (
	"fmt"
	"net/url"
)

func main() {
	data, _ := url.Parse("https://lco.dev/")
	fmt.Println(data.Host)
	fmt.Println(data.Port())
	fmt.Println(data.Scheme)
	fmt.Println(data.RawQuery)

}

package main

import (
	"fmt"
	"net/http"
	"sync"
)

var wg sync.WaitGroup

func GetStatusCode(url string) {
	defer wg.Done()
	resp, err := http.Get(url)
	if err != nil {
		fmt.Printf("%v is down\n", err)
	} else {
		fmt.Printf("status code of %v is %v\n", url, resp.StatusCode)
	}
	defer resp.Body.Close()
}

func main() {
	websites := []string{
		"https://udemy.com",
		"https://youtube.com",
		"https://yahoo.com",
		"https://amazon.com",
	}

	for _, web := range websites {
		go GetStatusCode(web)
		wg.Add(1)
	}

	fmt.Println("Waiting for goroutines to complete")
	wg.Wait()
	fmt.Println("All Jobs are Done!!")
}

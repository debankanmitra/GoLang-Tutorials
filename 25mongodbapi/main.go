package main

import (
	"fmt"
	"log"
	"mdb/router"
	"net/http"
)

func main() {
	r := router.Route()

	fmt.Println("SERVER IS STARTED")
	log.Fatal(http.ListenAndServe(":4000", r))
}

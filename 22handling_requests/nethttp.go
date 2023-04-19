// Handling equests and responses using standard library

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	// passing function inside function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>THIS IS HOME PAGE</h1>"))
	})

	// using handler function separately
	ProductsHandler := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "<h1>THIS IS PRODUCT PAGE</h1>\n")
	}
	http.HandleFunc("/products", ProductsHandler)

	// listening on port 3000
	log.Fatal(http.ListenAndServe(":3000", nil))
}

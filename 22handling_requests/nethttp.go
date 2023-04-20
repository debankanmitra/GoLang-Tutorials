// Handling equests and responses using standard library

package main

import (
	"io"
	"log"
	"net/http"
)

func main() {

	// passing controller inside function
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>THIS IS HOME PAGE</h1>"))
	})

	// using controller function separately
	ProductsHandler := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "<h1>THIS IS PRODUCT PAGE</h1>\n")
	}
	http.HandleFunc("/products", ProductsHandler)

	// listening on port 3000
	log.Fatal(http.ListenAndServe(":3000", nil))
}

/*
* request is a pointer because request is
* potentially quite large and complex with differnt types of data
* it is more efficient to pass a pointer to the request rather than a copy of the request.
* Passing a copy of the request would require copying all of its data, which could be quite expensive in terms of memory and processing time.
 */

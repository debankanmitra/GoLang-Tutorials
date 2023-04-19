// Handling equests and responses using gorilla library

package main

import (
	"io"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	r := mux.NewRouter()

	// passing function inside function
	r.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("<h1>THIS IS HOME PAGE FROM GORILLA</h1>"))
	})

	// using handler function separately
	r.HandleFunc("/products", ProductsHandler).Methods("GET")

	// listening on port 3000
	log.Fatal(http.ListenAndServe(":4000", r))
}

func ProductsHandler(w http.ResponseWriter, _ *http.Request) {
	io.WriteString(w, "<h1>THIS IS PRODUCT PAGE FROM GORILLA</h1>\n")
}

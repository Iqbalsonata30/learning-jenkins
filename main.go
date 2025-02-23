package main

import (
	"fmt"
	"log"
	"net/http"
)

func handler(w http.ResponseWriter, req *http.Request) {
	fmt.Fprintf(w, "Hello World\n")
}

func main() {
	http.HandleFunc("/", handler)
	log.Println("Server starting on :8080")
	log.Fatal(http.ListenAndServe(":8000", nil))
}

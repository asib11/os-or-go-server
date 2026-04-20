package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, World!")
}

func aboutHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "I'm Asib and I am learning Go!")
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/hello", helloHandler)	
	mux.HandleFunc("/about", aboutHandler)

	fmt.Println("Server is running on http://localhost:3000")
	err := http.ListenAndServe(":3000", mux)
	if err != nil {
		fmt.Println("Error starting server:", err)
	}
}

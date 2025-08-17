package main

import (
	"fmt"
	"net/http"
)

func sayHello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello, CI/CD World!")
}

func main() {
	http.HandleFunc("/", sayHello)
	fmt.Println("Server starting on port 8080...")
	http.ListenAndServe(":8080", nil)
}

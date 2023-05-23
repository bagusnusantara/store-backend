package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/hello", helloServer)
	fmt.Println("Starting server on port: 8080")
	http.ListenAndServe(":8080", nil)
}

func helloServer(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello World! HEHEH")
}

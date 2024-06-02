package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellohandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello World!")
}

func main() {
	http.HandleFunc("/hello", hellohandler)
	fmt.Print("hello the app is starting")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

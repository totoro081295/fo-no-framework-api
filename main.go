package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", getHelloWorld)
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func getHelloWorld(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello world!")
}

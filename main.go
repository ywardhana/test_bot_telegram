package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {

	port := ":8082"

	fmt.Printf("Kingsman is active...\n")
	fmt.Printf("Listening to port: 8082...\n")
	log.Fatal(http.ListenAndServe(port, router))
}

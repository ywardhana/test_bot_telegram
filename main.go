package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/bukalapak/kingsman/app/kingsman"
)

func main() {

	port := ":8082"

	router := kingsman.NewKingsman()

	fmt.Printf("Kingsman is active...\n")
	fmt.Printf("Listening to port: 8082...\n")
	log.Fatal(http.ListenAndServe(port, router))
}

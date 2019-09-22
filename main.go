package main

import (
	"log"
	"net/http"
	"os"
)



func main() {
	http.HandleFunc("/", listHandler)

	port := os.Args[1]

	log.Fatal(http.ListenAndServe(":"+port, nil))
}

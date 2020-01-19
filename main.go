package main

import (
	"log"
	"net/http"
	"os"
)

func initRouter() {
	http.HandleFunc("/", listHandler)
	http.HandleFunc("/create", newHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
}

func initFunc() {
	initRouter()

	err := initFileDao()
	if err != nil {
		panic(err)
	}

}

func main() {

	initFunc()

	// Port
	port := os.Args[1]

	// Start web service
	log.Println("start")
	log.Fatal(http.ListenAndServe(":"+port, nil))
}

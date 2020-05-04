package main

import (
	"flag"
	"fmt"
	"log"
	"net/http"
)

var (
	port    int
	address string
)

func initRouter() {
	http.HandleFunc("/create", newHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/", listHandler)
}

func iniFlag() {
	flag.IntVar(&port, "port", 8080, "server port")
	flag.StringVar(&address, "address", "127.0.0.1", "server address")

	rands := NewRandString()

	NewPostSugar = rands(10)

	SugarCounter = 0

	flag.Parse()

}

func initLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)
}

func initFunc() {

	initLog()

	iniFlag()

	initRouter()

	err := initFileDao()
	if err != nil {
		panic(err)
	}

}

func main() {

	initFunc()

	// Start web service
	addr := fmt.Sprintf("%s:%d", address, port)
	log.Printf("start at http://%s/\n", addr)
	log.Printf("New Sugar at http://%s/create?p=%s\n", addr, NewPostSugar)
	log.Fatal(http.ListenAndServe(addr, nil))
}

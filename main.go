package main

import (
	"flag"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
	"path"
)

var (
	port    int
	address string
)

func initStatic() {

	http.HandleFunc("/favicon.ico", func(w http.ResponseWriter, r *http.Request){
		http.ServeFile(w,r,"./templates/favicon.ico")
	})


	fs := http.FileServer(http.Dir("./static"))
	http.Handle("/static/", http.StripPrefix("/static/", fs))
}

func initRouter() {

	initStatic()

	http.HandleFunc("/create", newHandler)
	http.HandleFunc("/update", updateHandler)
	http.HandleFunc("/delete", deleteHandler)
	http.HandleFunc("/", listHandler)
}

func iniFlag() {
	flag.IntVar(&port, "port", 8080, "server port")
	flag.StringVar(&address, "address", "127.0.0.1", "server address")
	flag.StringVar(&LogsDir, "log", "logs/", "logs file")

	rands := NewRandString()

	NewPostSugar = rands(10)

	SugarCounter = 0

	flag.Parse()

}

func initLog() {
	log.SetFlags(log.Ldate | log.Ltime | log.Lshortfile)

	f, err := os.OpenFile(path.Join(LogsDir, "timeline.log"),
		os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
	if err != nil {
		log.Println(err)
	}

	log.SetOutput(io.MultiWriter(f, os.Stdout))
}

func initFunc() {

	iniFlag()

	initLog()

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

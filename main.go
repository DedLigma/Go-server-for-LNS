package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
)

func main() {

	dataCreator()
	go startServer()
	startObserver()

}

func startServer() {

	http.HandleFunc("/", mainPage)

	port := ":3001"
	fmt.Println("Server listen on port: ", port)

	err := http.ListenAndServe(port, nil)
	if err != nil {
		log.Fatal("ListenAndServer", err)
	}

}

func startObserver() {
	var pathWatcher Publisher = &PathWatcher{
		rootPath: "./srnslog",
	}
	var pathIndexer Subscriber = &PathIndexer{}
	pathWatcher.register(&pathIndexer)

	var pathFileMD5 Subscriber = &PathFileMD5{}
	pathWatcher.register(&pathFileMD5)

	pathWatcher.observe()

}

func mainPage(w http.ResponseWriter, r *http.Request) {

	js, _ := json.Marshal(data)

	w.Write(js)
}

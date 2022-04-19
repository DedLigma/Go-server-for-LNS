package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	// "os"
	// "strings"
)

var data struct {
	Solv struct {
		Start_timer int    `json:"start_timer"`
		Time_data   string `json:"time_data"`
		Timer       string `json:"timer"`
		Coordinates struct {
			Latitude int `json:"latitude"`
			Longtude int `json:"longtude"`
		} `json:"coordinates"`
		Height        int    `json:"height"`
		Svs           string `json:"svs"`
		Res           int    `json:"res"`
		Modi_ns       string `json:"modi_ns"`
		Modi_ms       string `json:"modi_ms"`
		Pdop          int    `json:"pdop"`
		E_distance    int    `json:"e_distance"`
		Appar_counter int    `json:"appar_counter"`
	}
	Lct struct {
		Number      []int `json:"number"`
		Time_data   []int `json:"time_data"`
		Timer       []int `json:"timer"`
		Coordinates struct {
			Latitude []int `json:"latitude"`
			Longtude []int `json:"longtude"`
		} `json:"coordinates"`
		Height []int   `json:"heights"`
		Resids [][]int `json:"resids"`
	}
}

func main() {

	// pathToFile := "srns_0x0003.txt"
	// dat, _ := os.ReadFile(pathToFile)       //считывание файла
	// arr := strings.Split(string(dat), "\n") //заносим все строки в массив
	// lastElemntsArr := arr[len(arr)-451 : len(arr)-1]

	// data.Solv.start_timer = 123123
	// fmt.Println(string(data))

	startServer()
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

func mainPage(w http.ResponseWriter, r *http.Request) {

	js, _ := json.Marshal(data)

	w.Write(js)
}

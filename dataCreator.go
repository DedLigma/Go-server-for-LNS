package main

import (
	"os"
	"strings"
)

func dataCreator() {
	pathToFile := "./srnslog/srns_0x0003.txt"
	dat, _ := os.ReadFile(pathToFile)       //считывание файла\
	arr := strings.Split(string(dat), "\n") //заносим все строки в массив
	lastElemntsArr := arr[len(arr)-451 : len(arr)-1]
	data.Solv.Time_data = lastElemntsArr[449]
	// fmt.Println(lastElemntsArr[449])
}

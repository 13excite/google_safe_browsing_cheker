package main

import (
	"fmt"
	"io/ioutil"
)

func getRequestJson(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error open file %s", err)
	}
	fmt.Printf(string(data))
	return data
}

func main() {
	const myReq string = "/tmp/json.txt"
	getRequestJson(myReq)
}

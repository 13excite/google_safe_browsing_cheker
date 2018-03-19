package main

import (
	"bufio"
	"fmt"
	"net/http"
	"os"
	//"encoding/json"

	"bytes"
	"io/ioutil"
)

func getURL(filePath string) string {
	file, err := os.Open(filePath)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	var lines []string
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		lines = append(lines, scanner.Text())
	}
	gsfKey := lines[0]
	url := "http://ddd.ru/"
	return url + gsfKey
}

func sendRequest(requestUrl string) {
	var jsonStr = []byte(`{"title":"test"}`)
	req, err := http.NewRequest("POST", requestUrl, bytes.NewBuffer(jsonStr))
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response status: ", resp.Status)
	fmt.Println("response header:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response body:", string(body))
}

func main() {
	const URL string = "http://ya.ru"
	//fmt.Printf(getURL("/tmp/test.txt"))
	fmt.Println("bla bla")
	sendRequest(URL)
}

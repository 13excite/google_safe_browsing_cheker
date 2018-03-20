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

func getURL(filePath, shortRequestURL string) string {
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
	return shortRequestURL + gsfKey
}

func sendRequest(requestURL string) {
	var jsonStr = []byte(`{
        "client": {
                    "clientId":      "myproject",
                    "clientVersion": "1.5.2"
                },
        "threatInfo": {
                    "threatTypes":      ["MALWARE", "SOCIAL_ENGINEERING", "POTENTIALLY_HARMFUL_APPLICATION", "UNWANTED_SOFTWARE"],
                    "platformTypes":    ["ANY_PLATFORM"],
                    "threatEntryTypes": ["URL"],
                    "threatEntries": [
                                {"url": "http://malware.testing.google.test/testing/malware/"},
                                {"url": "http://ianfette.org"},
                            ]
                }
}`)
	req, err := http.NewRequest("POST", requestURL, bytes.NewBuffer(jsonStr))
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
	const URL string = "https://safebrowsing.googleapis.com/v4/threatMatches:find?key="
	fullURL := getURL("/tmp/google.key", URL)
	sendRequest(fullURL)
}

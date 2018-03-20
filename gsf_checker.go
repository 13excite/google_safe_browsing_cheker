package main

import (
	"bufio"
	"flag"
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
	filePath := flag.String("key", "", "usage --key /foo/bar.key")
	flag.Parse()

	key := *filePath
	// check required flag
	if key == "" {
		flag.Usage()
		os.Exit(1)
	}
	const URL string = "https://safebrowsing.googleapis.com/v4/threatMatches:find?key="
	//fullURL := getURL("/tmp/google_sb.key", URL)
	sendRequest(getURL(key, URL))
}

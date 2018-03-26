package main

import (
	"bufio"
	"bytes"
	"encoding/json"
	"flag"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"os"
	"time"
)

type Threat struct {
	Url string `json:"url"`
}

type APIResponse struct {
	Matches []Matches `json:"matches"`
}
type Matches struct {
	ThreatType      string  `json:"threatType"`
	PlatformType    string  `json:"platformType"`
	Threat          *Threat `json:"threat"`
	CacheDuration   string  `json:"cacheDuration"`
	ThreatEntryType string  `json:"threatEntryType"`
}

func getJsonOfRequest(filePath string) []byte {
	data, err := ioutil.ReadFile(filePath)
	if err != nil {
		fmt.Printf("Error open file: %s", err)
		os.Exit(1)
	}
	return data
}
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

func sendRequest(requestURL string) []byte {
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

	client := &http.Client{Timeout: 4 * time.Second}
	resp, err := client.Do(req)
	if err != nil {
		panic(err)
	}
	defer resp.Body.Close()
	fmt.Println("response status: ", resp.Status)
	fmt.Println("response header:", resp.Header)

	body, _ := ioutil.ReadAll(resp.Body)
	fmt.Println("response body:", string(body))
	return body
}

func parseJson(responseData []byte) {
	response := bytes.NewReader(responseData)
	decoder := json.NewDecoder(response)
	val := &APIResponse{}
	err := decoder.Decode(val)
	if err != nil {
		log.Fatal(err)
	}
	for _, s := range val.Matches {
		fmt.Println(s.Threat.Url)
	}
}

func main() {
	const URL string = "https://safebrowsing.googleapis.com/v4/threatMatches:find?key="

	keyPath := flag.String("key", "", "usage --key /foo/bar.key")
	flag.Parse()

	key := *keyPath
	// check required flag
	if key == "" {
		flag.Usage()
		os.Exit(1)
	}

	parseJson(sendRequest(getURL(key, URL)))
	//sendRequest(getURL(key, URL))
}

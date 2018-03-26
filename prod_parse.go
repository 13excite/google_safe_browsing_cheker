package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"os"
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
func main() {
	data := getJsonOfRequest("/tmp/json.txt")
	r := bytes.NewReader(data)
	decoder := json.NewDecoder(r)

	val := &APIResponse{}
	err := decoder.Decode(val)
	if err != nil {
		log.Fatal(err)
	}
	//fmt.Println(val)
	for _, s := range val.Matches {
		fmt.Println(s.Threat.Url)

	}

}

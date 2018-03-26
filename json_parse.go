package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"log"
)

type Example struct {
	Type    string   `json:"type,omitempty"`
	Subsets []Subset `json:"subsets,omitempty"`
}

type Subset struct {
	Addresses []Address `json:"addresses,omitempty"`
}

type Address struct {
	IP string `json:"IP,omitempty"`
}

func main() {

	m := []byte(`{"type":"example","data": {"name": "abc","labels": {"key": "value"}},"subsets": [{"addresses": [{"ip": "192.168.103.178"}],"ports": [{"port": 80}]}]}`)

	r := bytes.NewReader(m)
	decoder := json.NewDecoder(r)

	val := &Example{}
	err := decoder.Decode(val)

	if err != nil {
		log.Fatal(err)
	}

	// If you want to read a response body
	// decoder := json.NewDecoder(res.Body)
	// err := decoder.Decode(val)

	// Subsets is a slice so you must loop over it
	for _, s := range val.Subsets {
		// within Subsets, address is also a slice
		// then you can access each IP from type Address
		for _, a := range s.Addresses {
			fmt.Println(a.IP)
		}
	}

}

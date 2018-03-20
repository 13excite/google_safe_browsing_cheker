package main

import (
	"flag"
	"fmt"
)

func main() {
	testStr := flag.String("world", "ddd", "zzzz")
	numPort := flag.Int("int", 42, "an int")

	flag.Parse()

	fmt.Println("srt: ", *testStr)
	fmt.Println("int: ", *numPort)
}

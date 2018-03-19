package main

import (
	"bufio"
	"fmt"
	"os"
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
func main() {
	fmt.Printf(getURL("/tmp/test.txt"))
	fmt.Println("bla bla")
}
